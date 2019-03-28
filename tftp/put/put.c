
#include <assert.h>
#include <sys/types.h>          /* See NOTES */
#include <sys/socket.h>
#include <string.h>
#include <stdint.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <stdio.h>


#define PORT    69
#define HOST    "192.168.90.162"

#define FILENAME    "put"
#define MODE        "octet"

#define RDQ     1
#define WRQ     2
#define DATA    3
#define ACK     4
#define ERR     5


struct Request {
    const char *filename;
    const char *mode;
    uint16_t method;
};
typedef struct Request Request;

size_t build_req_packet(Request *p, void *buffer) {
    char *buf = (char *)buffer;
    size_t len;

    uint16_t opcode = htons(p->method);
    memcpy(buf, &opcode, sizeof(opcode));
    buf += 2;

    len = strlen(p->filename) + 1;
    memcpy(buf, p->filename, len);
    buf += len;

    len = strlen(p->mode) + 1;
    memcpy(buf, p->mode, len);
    buf += len;
    return (void *)buf - buffer;
}


struct ACKPacket {
    uint16_t opcode;
    uint16_t blocknr;
};


struct DATAPacket {
    uint16_t opcode;
    uint16_t blocknr;
    char data[0];
};

int doack(void *buf, size_t len, uint16_t blocknr) {
    assert(len == sizeof(struct ACKPacket));
    struct ACKPacket *p = (struct ACKPacket *)buf;
    assert(p->opcode == htons(ACK));
    uint16_t nr = ntohs(p->blocknr);
    assert(nr <= blocknr);
    if (nr == blocknr) return 0;
    return 1;
}

// 11 14
// 44 26

void doput(int sock, struct sockaddr_in peer, socklen_t addrlen, const char *filename) {

    FILE *fp = fopen(filename, "rb");
    fseek(fp, 0L, SEEK_END);
    size_t flen = ftell(fp);
    uint16_t nr = flen / 512 + 1;
    fseek(fp, 0L, SEEK_SET);

    char rbuf[1024];
    char buf[1024];
    struct DATAPacket *p = (struct DATAPacket *)buf;

    struct sockaddr_in cli_addr;

    p->opcode = htons(DATA);
    for (uint16_t i = 1; i <= nr; i++) {
        size_t rdlen = fread(p->data, 1, 512, fp);
        p->blocknr = htons(i);

        sendto(sock, buf, rdlen+sizeof(struct DATAPacket) , 0, (struct sockaddr *)&peer, sizeof(peer));

        for (;;) {
            ssize_t r = recvfrom(sock, rbuf, sizeof(rbuf), 0, (struct sockaddr *)&cli_addr, &addrlen);
            if (doack(rbuf, r, i) == 0) 
                break;
        }
    }
    fclose(fp);
}

int main() {
    struct sockaddr_in srv_addr, cli_addr;
    socklen_t addrlen = sizeof(cli_addr);
    int sock = socket(AF_INET, SOCK_DGRAM, 0);

    memset(&srv_addr, 0, sizeof(srv_addr));
    srv_addr.sin_family = AF_INET;
    srv_addr.sin_port = htons(PORT);
    srv_addr.sin_addr.s_addr = inet_addr(HOST);

    Request p;
    p.method = WRQ;
    p.filename = FILENAME;
    p.mode = MODE;

    char buf[1024];

    size_t len = build_req_packet(&p, buf);

    sendto(sock, buf, len, 0, (struct sockaddr *)&srv_addr, sizeof(srv_addr));
    ssize_t r = recvfrom(sock, buf, sizeof(buf), 0, (struct sockaddr *)&cli_addr, &addrlen);
    doack(buf, r, 0);

    doput(sock, cli_addr, addrlen, p.filename);
    return 0;
}


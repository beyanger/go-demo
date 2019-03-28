
package main

import (
    "bytes"
    "strconv"
    "net"
    "fmt"
    "encoding/binary"
)

var (
    fileName = "put.go"
    host = "192.168.90.162"
    port = 69
)


const (
    RDQ = iota+1
    WRQ
    DATA
    ACK
    ERR
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}


type Packet interface {
    Type() uint16
    Bytes() []byte
}


type WRQPacket struct {
    Filename    string
    Mode        string
}

func (p *WRQPacket) Type() uint16 {
    return WRQ
}

func (p *WRQPacket) Bytes() []byte {
    buf := new(bytes.Buffer)
    opcode := make([]byte, 2)
    binary.BigEndian.PutUint16(opcode, p.Type())
    buf.Write(opcode)
    buf.WriteString(p.Filename)
    buf.WriteByte(0)
    buf.WriteString(p.Mode)
    buf.WriteByte(0)
    return buf.Bytes()
}




func main() {

    addr, err := net.ResolveUDPAddr("udp", host + ":" + strconv.Itoa(port))
    checkErr(err)

    conn, err := net.ListenUDP("udp", nil, addr)
    checkErr(err)

    packet := &WRQPacket{fileName, "octet"}

    b := packet.Bytes()

    _, err = conn.Write(b)

    checkErr(err)


    data := make([]byte, 1024)
    _, remoteAddr, err := conn.ReadFromUDP(data)

    fmt.Println(remoteAddr, err)



    conn.Close()
}


package main

import (
    "encoding/binary"
    "bytes"
    "os"
    "fmt"
    "strings"
    "bufio"
)



var helpstr =
`Commands may be abbreviated.  Commands are:

connect 	connect to remote tftp
mode    	set file transfer mode
put     	send file
get     	receive file
quit    	exit tftp
verbose 	toggle verbose mode
trace   	toggle packet tracing
status  	show current status
binary  	set mode to octet
ascii   	set mode to netascii
rexmt   	set per-packet retransmission timeout
timeout 	set total retransmission timeout
?       	print help information
`

var command = [...]string{
    "connect",
    "mode",
    "put",
    "get",
    "quit",
    "verbose",
    "trace",
    "status",
    "binary",
    "ascii",
    "rexmt",
    "timeout",
    }

const (
        invalid = "?Invalid command"
        ambiguous = "?Ambiguous command"
        header = "tftp> "
        connHeader = "(to) "
        putHeader = "(file) "
)

var (
    server = ""
    mode = "netascii"
)


var reader = bufio.NewReader(os.Stdin)

func doConnect(param []string) {
    if len(param) > 0 {
        server = param[0]
        return
    }

    fmt.Print(connHeader)
    input, _, err := reader.ReadLine()
    if err == nil {
        server = string(input)
    }
}

func doMode() {
    if mode == "binary" {
        fmt.Println("Using octet mode to transfer files.")
    } else {
        fmt.Println("Using netascii mode to transfer files.")
    }
}


func doBinary() {
    mode = "binary"
}

func doAscii() {
    mode = "ascii"
}

func checkHost(server string) bool {
    if server == "" {
        fmt.Println("No target machine specified.")
        return false
    }
    return true
}


type RRQPacket struct {
    Filename    string
    Seq         uint16
}


func (p *RRQPacket) Type() uint16 {
    return RRQ
}

func (p *RRQPacket) Bytes() []byte {
    buf := new(bytes.Buffer)
    opcode := make([]byte, 2)
    binary.BigEndian.PutUint16(opcode, p.Type())
    buf.Write(opcode)
    buf.WriteString(p.Filename)
    return buf.Bytes()
}


type DATAPacket struct {
    Data        []byte
    Seq         uint16
}

type ACKPacket struct {
    Seq         uint16
}

func (p *ACKPacket) Type() uint16 {
    return ACK
}

func doPutFile(file string) error {
    return nil
}


func doPut(param []string) {

    file := param

    if len(file) == 0 {
        fmt.Print(putHeader)
        input, _, err := reader.ReadLine()
        if err == nil {
            return
        }
        file = strings.Split(string(input), " ")
    }

    if !checkHost(server) {
        return
    }

    for _, f := range file {
        doPutFile(f)
    }
}

func doGet(param []string) {
    if !checkHost(server) {
        return
    }
}

func doCommand(cmd string, param []string) {
    switch cmd {
    case "connect":
        doConnect(param)
    case "put":
        doGet(param)
    case "get":
        doPut(param)
    case "mode":
        doMode()
    case "binary":
        doBinary()
    case "ascii":
        doAscii()
    }
}

func checkCommand(cmd string) bool {
    for _, c := range cmd {
        if c < 'a' || c > 'z' {
            return false
        }
    }
    return true
}

func main() {
    trie := NewTrie()
    for i := range command {
        trie.Insert(command[i])
    }
    for {
        fmt.Print(header)
        input, _, err := reader.ReadLine()
        if err != nil || len(input) == 0 {
            continue
        }
        cmdlist := strings.Split(string(input), " ")
        rd := strings.ToLower(cmdlist[0])
        if !checkCommand(rd) {
            continue
        }
        if rd == "?" {
            fmt.Print(helpstr)
            continue
        }

        cmd := trie.Search(rd)

        switch len(cmd) {
        case 0: fmt.Println(invalid)
        case 1:
            if cmd[0] == "quit" {
                fmt.Println("byte")
                return
            }
            doCommand(cmd[0], cmdlist[1:])
        default: fmt.Println(ambiguous)
        }
    }
}


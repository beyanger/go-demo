
package main
import (
    "fmt"
    "net"
    "io"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
/*
    io.ReadWriterCloser 是一个接口，接口定义了三个方法 Read, Write, Close
    这里不关心 conn 到底是个什么类型，只要实现了上述三个方法就行...
*/

// origin: *TCPConn
func echo(conn io.ReadWriteCloser) {
    defer conn.Close()

    buffer := make([]byte, 1024)

    for {
        n, err := conn.Read(buffer)

        if err != nil {
            fmt.Println("read error return: ", err)
            return
        }

        fmt.Print("recv:", string(buffer[:n]))

        n, err = conn.Write(buffer[:n])
        if err != nil {
            fmt.Println("send error return: ", err)
            return
        }
    }
}

func main() {
    laddr, err := net.ResolveTCPAddr("tcp", ":9090")
    checkErr(err)

    ln, err := net.ListenTCP("tcp", laddr)
    checkErr(err)

    for {
        conn, err := ln.AcceptTCP()
        checkErr(err)
        go echo(conn)
    }
}


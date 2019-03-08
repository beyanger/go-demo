
package main

import (
    "fmt"
    "time"
    "bufio"
    "sort"
    "encoding/binary"
    "net"
    "math/rand"
    "io"
    "strconv"
)

var prepared = make(chan bool, CNT)
var done = make(chan bool, CNT)
var starttime = time.Now()

func RandomSource(cnt int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < cnt; i++ {
            r := rand.Int()
            ch <- r
        }
        close(ch)
    }()
    return ch
}



func WriterSink(writer io.Writer, ch <-chan int) {
    buffer := make([]byte, 8)
    for v := range ch {
        binary.BigEndian.PutUint64(buffer, uint64(v))
        writer.Write(buffer)
    }
}

func NetworkSink(seq int, ch <-chan int) {
    addr := SERVER + ":"+strconv.Itoa(seq+PORT)
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        panic(err)
    }

    prepared <- true
    go func() {
        conn, err := ln.Accept()
        if err != nil {
            panic(err)
        }
        writer := bufio.NewWriter(conn)
        WriterSink(writer, ch)
        writer.Flush()
        conn.Close()
        ln.Close()
        done <- true
    }()
}

func DataSource(seq, cnt int) {
    ch := RandomSource(cnt)
    data := make([]int, 0)
    for n := range ch {
        data = append(data, n)
    }

    sort.Ints(data)

    fmt.Printf("Worker: %d sort done with: %v\n", seq, time.Now().Sub(starttime))
    sch := make(chan int)

    go func(){
        for _, n := range data {
            sch <- n
        }
        close(sch)
    }()
    NetworkSink(seq, sch)
}

func PrepareData() {
    for i := 0; i < CNT; i++ {
        go DataSource(i, SIZE)
    }
}

func main() {

    PrepareData()
    for i := 0; i < CNT; i++ {
        <-prepared
    }

    fmt.Printf("All workder prepared with: %v\n", time.Now().Sub(starttime))

    for i := 0; i < CNT; i++ {
        <-done
    }

    fmt.Printf("All workder done with: %v\n", time.Now().Sub(starttime))
}

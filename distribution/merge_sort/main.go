
package main

import (
    "fmt"
    "os"
    "time"
    "bufio"
    "sort"
    "encoding/binary"
    "net"
    "math/rand"
    "io"
    "strconv"
)

var wg = make(chan bool, CNT)
var starttime = time.Now()

const (
    CNT = 10
    SIZE = 1e3
    PORT = 8000
)

func RandomSource(cnt int) <-chan int {
    ch := make(chan int)

    go func() {
        for i := 0; i < cnt; i++ {
            r := rand.Int()
            fmt.Println(r)
            ch <- r
        }
        close(ch)
    }()
    return ch
}


func NetworkSource(seq, size int) <-chan int {
    ch := make(chan int)

    go func() {
        addr := ":" + strconv.Itoa(PORT + seq)
        conn, err := net.Dial("tcp", addr)

        if err != nil {
            panic(err)
        }

        reader := bufio.NewReader(conn)

        buffer := make([]byte, 8)

        for i := 0; i < size; i++ {
            n, err := reader.Read(buffer)
            if n > 0 {
                val := binary.BigEndian.Uint64(buffer)
                ch <-int(val)
            }
            if err != nil {
                break
            }
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
    addr := ":"+strconv.Itoa(seq+PORT)
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        panic(err)
    }

    wg <- true

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

func Merge(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    go func(){
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2

        for ok1 || ok2 {
            if ok1 && (!ok2 || v1 > v2) {
                out <-v1
                v1, ok1 = <-ch1
            } else {
                out <-v2
                v2, ok2 = <-ch2
            }
        }
        close(out)
    }()
    return out
}

func MergeN(data []<-chan int) <-chan int {
    ld := len(data)
    if ld == 1 {
        return data[0]
    }
    return Merge(MergeN(data[:ld/2]), MergeN(data[ld/2:]))
}

func main() {

    PrepareData()
    // wait all data prepared
    for i := 0; i < CNT; i++ {
        <-wg
    }

    fmt.Printf("All workder done with: %v\n", time.Now().Sub(starttime))

    data := make([]<-chan int, 0)
    for i := 0; i < CNT; i++ {
        dt := NetworkSource(i, SIZE)
        data = append(data, dt)
    }
    sn := MergeN(data)

    outfile := "out.file"
    file, err := os.OpenFile(outfile, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }

    writer := bufio.NewWriter(file)
    WriterSink(writer, sn)
    writer.Flush()
    file.Close()
    fmt.Printf("Merge sort done with: %v\n", time.Now().Sub(starttime))
}


package main

import (
    "fmt"
    "time"
    "os"
    "bufio"
    "sort"
    "encoding/binary"
    "net"
    "math/rand"
    "io"
    "strconv"
)


const (
    COUNT = 10
    SIZE = 10
    PORT = 8000
    SERVER = "127.0.0.1"
    OUTFILE = "out.txt"
)

var dataPrepared = make(chan struct{}, COUNT)
var starttime = time.Now()

func RandomSource(cnt int) <-chan int {
    ch := make(chan int)

    go func() {
        for i := 0; i < cnt; i++ {
            ch <- rand.Int()
        }
        close(ch)
    }()
    return ch
}


func NetworkSource(seq, size int) <-chan int {
    ch := make(chan int)
    go func() {
        addr := SERVER + ":" + strconv.Itoa(PORT + seq)
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

func binaryWriterSink(writer io.Writer, ch <-chan int) {
    buffer := make([]byte, 8)
    for v := range ch {
        binary.BigEndian.PutUint64(buffer, uint64(v))
        writer.Write(buffer)
    }
}

func textWriterSink(writer io.Writer, ch <-chan int) {
    for v := range ch {
        tv := strconv.Itoa(v)
        writer.Write([]byte(tv))
        writer.Write([]byte{'\n'})
    }
}

func NetworkSink(seq int, ch <-chan int) {
    addr := SERVER+":"+strconv.Itoa(seq+PORT)
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Worker %d data prepared with %v\n", seq, time.Now().Sub(starttime))
    dataPrepared <- struct{}{}

    go func() {
        conn, err := ln.Accept()
        if err != nil {
            panic(err)
        }
        writer := bufio.NewWriter(conn)
        binaryWriterSink(writer, ch)
        writer.Flush()
        conn.Close()
        ln.Close()
    }()
}

func DataSource(seq, cnt int) {
    data := make([]int, 0, cnt)
    for i := 0; i < cnt; i++ {
        data = append(data, rand.Int())
    }

    sort.Ints(data)

    sch := make(chan int, 8)

    go func() {
        for i := range data {
            sch <- data[i]
        }
        close(sch)
    }()
    NetworkSink(seq, sch)
}

func PrepareData() {
    for i := 0; i < COUNT; i++ {
        go DataSource(i, SIZE)
    }
}

func Merge(ch1, ch2 <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2

        for ok1 || ok2 {
            if !ok2 || (ok1 && v1 <= v2) {
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

    rand.Seed(time.Now().UnixNano())

    PrepareData()
    // wait all data prepared
    for i := 0; i < COUNT; i++ {
        <-dataPrepared
    }

    data := make([]<-chan int, 0)
    for i := 0; i < COUNT; i++ {
        dt := NetworkSource(i, SIZE)
        data = append(data, dt)
    }
    sn := MergeN(data)

    file, err := os.OpenFile(OUTFILE, os.O_RDWR|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }

    writer := bufio.NewWriter(file)
    textWriterSink(writer, sn)
    writer.Flush()
    file.Close()
    fmt.Printf("Merge sort done with: %v\n", time.Now().Sub(starttime))
}


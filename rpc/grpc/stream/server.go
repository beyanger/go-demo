
package main

import (
    "time"
    "sync"
    //"context"
    "strconv"
    "net"
    "fmt"
    srv "com/beyanger/service"
    grpc "google.golang.org/grpc"
)

const (
    PORT = 2444
    CNT = 3
)

type StreamServer struct {

}

func (s *StreamServer) GetStream(req *srv.StreamReqData, res srv.Greeter_GetStreamServer) error {
    fmt.Println("Get Stream: ", req)
    for i := 0; i < CNT; i++ {
        if err := res.Send(&srv.StreamResData{Data:fmt.Sprintf("server time: %v",time.Now().Unix())}); err != nil {
            fmt.Println("send error: ", err)
            break
        }
        time.Sleep(time.Second)
    }
    return nil
}


func (s *StreamServer) PutStream(req srv.Greeter_PutStreamServer) error {
    for {
        // SendAndClose
        if t, err := req.Recv(); err == nil {
            fmt.Println("put stream: ", t)
        } else {
            fmt.Println("break, err :", err)
            break
        }
    }

    if err := req.SendAndClose(&srv.StreamResData{Data:"hahahah"}); err != nil {
        fmt.Println(err)
    }

    return nil
}

func (s *StreamServer) AllStream(all srv.Greeter_AllStreamServer) error {
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        for {
            if data, err := all.Recv(); err == nil {
                fmt.Println("recv data: ", data)
            } else {
                fmt.Println("recv error: ", err)
                break
            }
        }
        wg.Done()
    }()

    go func() {
        for i := 0; i < CNT; i++ {
            if err := all.Send(&srv.StreamResData{Data:fmt.Sprintf("server time: %v",time.Now().Unix())}); err != nil {
                fmt.Println("send err: ", err)
                break
            }
            time.Sleep(time.Second)
        }
        wg.Done()
    }()
    wg.Wait()
    return nil
}

func main() {
    ln, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
    if err != nil {
        panic(err)
    }

    s := grpc.NewServer()
    srv.RegisterGreeterServer(s, &StreamServer{})
    s.Serve(ln)
}

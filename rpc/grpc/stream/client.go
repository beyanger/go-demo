
package main
import (
    "time"
    "context"
   "fmt"
    "strconv"
    srv "com/beyanger/service"
    grpc "google.golang.org/grpc"
)

const (
    PORT = 2444
    CNT = 3
)

func main() {
    conn ,err := grpc.Dial("192.168.90.162:"+strconv.Itoa(PORT), grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    c := srv.NewGreeterClient(conn)

    reqData := &srv.StreamReqData{Data:"aaa"}

    res, err := c.GetStream(context.Background(), reqData)
    if err != nil {
        panic(err)
    }
    for {
        if data, err := res.Recv(); err == nil {
            fmt.Println("server stream: ", data.Data)
        } else {
            fmt.Println("recv error: ", err)
            break
        }
    }

    put, err := c.PutStream(context.Background())
    if err != nil {
        panic(err)
    }
    for i := 0; i < CNT; i++ {
        content := fmt.Sprintf("client time: %v", time.Now())
        data := srv.StreamReqData{Data:content}
        if err := put.Send(&data); err != nil {
            panic(err)
        }
        time.Sleep(time.Second)
    }

    if data, err := put.CloseAndRecv(); err == nil {
        //
        fmt.Println("close recv data ", data)
    } else {
        fmt.Println("close err: ", err)
    }
}


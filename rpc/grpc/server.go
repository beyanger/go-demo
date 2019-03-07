
package main
import (
    "fmt"
    "context"
    "net"
    srv "com/beyanger/service"
    grpc "google.golang.org/grpc"
)


type dataCenter struct {

}

var age int32 = 22

func (dc *dataCenter) GetUserInfo(ctx context.Context, req *srv.UserRequest) (resp *srv.UserResponse, err error) {

    name := req.GetName()
    if name == "yang" {
        age++
        resp = &srv.UserResponse{
            Id:1,
            Name:"yang shuangyi",
            Age: age,
            Title: []string{"shabi", "danteng", "hahaha"},
        }
    }
    err = nil
    return
}

func main() {
    fmt.Println(grpc.Version)
    fmt.Println(grpc.ErrClientConnClosing)
    l, err := net.Listen("tcp", ":2333")
    if err != nil {
        panic(err)
    }
    server := grpc.NewServer()
    dc := dataCenter{}
    srv.RegisterUserInfoServiceServer(server, &dc)
    server.Serve(l)
}


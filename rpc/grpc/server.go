
package main
import (
    "context"
    "net"
    srv "com/beyanger/service"
    grpc "google.golang.org/grpc"
)


type dataCenter struct {
    age int32
}

func (dc *dataCenter) GetUserInfo(ctx context.Context, req *srv.UserRequest) (resp *srv.UserResponse, err error) {

    name := req.GetName()
    if name == "yang" {
        dc.age++
        resp = &srv.UserResponse{
            Id:1,
            Name:"yang shuangyi",
            Age: dc.age,
            Title: []string{"shabi", "danteng", "hahaha"},
        }
    }
    err = nil
    return
}

func main() {
    l, err := net.Listen("tcp", ":2333")
    if err != nil {
        panic(err)
    }
    server := grpc.NewServer()
    dc := dataCenter{age:22}
    srv.RegisterUserInfoServiceServer(server, &dc)
    server.Serve(l)
}


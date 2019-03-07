
package main
import (
    "context"
    "fmt"
    srv "com/beyanger/service"
    grpc "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:2333", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    client := srv.NewUserInfoServiceClient(conn)
    req := &srv.UserRequest{Name:"yang"}

    resp, err := client.GetUserInfo(context.Background(), req)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp)
}



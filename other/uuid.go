
package main

import (
    "github.com/satori/go.uuid"
    "fmt"
)

func main() {
    u1, err := uuid.NewV4()
    if err != nil {
        panic(err)
    }

    fmt.Printf("%T, %v\n", u1, u1)
}

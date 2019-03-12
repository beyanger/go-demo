
package main

import (
    "fmt"
    "reflect"
)


func main() {
    at := "abcdefg"
    fmt.Printf("type: %T, %v\n", reflect.TypeOf(at), reflect.TypeOf(at))
    fmt.Printf("val: %T, %v\n", reflect.ValueOf(at), reflect.ValueOf(at))
}

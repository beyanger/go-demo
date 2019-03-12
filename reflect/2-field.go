
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id      int
    Name    string
    Age     int
}


func (u User) ReflectCallFunc() {
    fmt.Printf("%s reflectCallFunc\n", u.Name)
}
func (u User) ReflectCallFunc2() {
    fmt.Printf("%s reflectCallFunc\n", u.Name)
}



func main() {
    user := User{1, "yang", 22}
    doFieldAndMethod(user)
}

func doFieldAndMethod(input interface{}) {
    getType := reflect.TypeOf(input)
    getVal := reflect.ValueOf(input)

    for i := 0; i < getType.NumField(); i++ {
        field := getType.Field(i)
        value := getVal.Field(i)
        fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
    }

    for i := 0; i < getType.NumMethod(); i++ {
        m := getType.Method(i)
        fmt.Printf("%s: %v\n", m.Name, m.Type)
    }
}

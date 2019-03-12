
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

func (u User) ReflectCallWithParam(name string, age int) {
    fmt.Printf("---------%s reflect call with param func  %v\n", name, age)
}



func main() {
    user := User{1, "yang", 22}
    doFieldAndMethod(user)
}

func doFieldAndMethod(input interface{}) {
    getType := reflect.TypeOf(input)
    getVal := reflect.ValueOf(input)

    methodVal := getVal.MethodByName("ReflectCallWithParam")
    args := []reflect.Value{reflect.ValueOf("yang"), reflect.ValueOf(11)}
    methodVal.Call(args)

    for i := 0; i < getType.NumMethod(); i++ {
        m := getType.Method(i)
        fmt.Printf("%s: %v\n", m.Name, m.Type)
    }
}


package main

import (
    "fmt"
)

type Person struct {
    name    string
    age     int
}

func (p *Person) SetAge(age int) *Person {
    p.age = age
    return p
}

func (p *Person) SetName(name string) *Person {
    p.name = name
    return p
}

func (p *Person) Print() *Person {
    fmt.Println("name: ", p.name, " age: ", p.age)
    return p
}

func main() {
    p := &Person{"yang", 20}
    q.Print().SetAge(30).Print().SetName("hahah").Print()
}


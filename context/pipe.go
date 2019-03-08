
package main

import (
    "fmt"
    "sync"
)

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup

    output := func(c <-chan int) {
        for n := range c {
            select {
            case out <- n:
            case <- done:
                return
            }
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }
    go func() {
        wg.Done()
    }
    return out
}


func gen(num ...int) <-chan int {
    ch := make(chan int)
    return ch
}

func sq(in <-chan int) <-chan int {
    return in
}


func main() {
    in := gen(2, 3)
    c1 := sq(in)
    c2 := sq(in)

    done := make(chan struct{}, 2)
    out := merge(done, c1, c2)

    fmt.Println(<-out)
    done <- struct{}{}
    done <- struct{}{}
}

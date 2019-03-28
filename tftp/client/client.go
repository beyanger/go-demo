

package main

import (
    "fmt"
    "time"
    "os"
)

func doPutFile(file string) {

}

func main() {
    if len(os.Args) < 2 {
        return
    }

    start := time.Now()
    doPutFile(os.Args[1])
    fmt.Println(time.Now().Sub(start))
}


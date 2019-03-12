
// addressable: pointer

package main
import (
    "fmt"
    "reflect"
)


func main() {
    var num float64 = 1.2345
    fmt.Println("old value is ", num)

    ptr := reflect.ValueOf(&num)
    // 如果上面ValueOf不是传递的指针，这里不可取
    newValue := ptr.Elem() // just ok when ptr is a pointer, will panic if not
    fmt.Println("type of pointer:", newValue.Type())
    fmt.Println("settability of pointer:", newValue.CanSet())

    if newValue.CanSet() {
        // newValue是原始对象的反射，只有原始对象才可以set
        newValue.SetFloat(77)
    }
    fmt.Println("new value is ", num)
}

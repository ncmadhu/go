package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    var a int
    var b int
    a = 5
    b = 6
    var c int = add(a, b)
    fmt.Println(c)
}

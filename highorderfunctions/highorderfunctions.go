package main

import "fmt"

type fn func(int) int

/* Higher order functions takes function as an argument or returns a function as its result */
func twice(f fn, x int) int {
    return f(f(x))
}

func add3(x int) int {
    return x + 3
}

func add() fn {
    return add3
}


func main() {
    var a int = 5
    var b int = twice(add3, a)
    fmt.Printf("%d\n", b)
    var c fn = add()
    fmt.Printf("%d\n", c(3))
}

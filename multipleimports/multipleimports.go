package main

import (
    "fmt"
    "strings"
)

func stringjoin(a []string) (string) {
    return strings.Join(a, " ")
}

func main() {
    var greetings = []string {"hello", "world!"} 
    var length int
    fmt.Printf("%s\n", stringjoin(greetings))
    fmt.Printf("%T\n", greetings)
    length = len(greetings[0])
    fmt.Printf("%d\n", length)
    fmt.Printf("%s\n", greetings[0])
}

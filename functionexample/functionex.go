package main

import "fmt"

func swap(a, b string) (string, string) {
    return b, a
}

func main() {
    var a = "Madhu"
    var b = "Chakravarthy"
    a, b = swap(a, b)
    fmt.Printf("lastname: %s\tfirstname: %s\n", a,b)
} 

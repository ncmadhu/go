package main

import (
	"fmt"
)

func SliceInit(x, y int) [][]uint8 {
	
	var a = make([][]uint8, x)
	for i := range a {
		a[i] = make([]uint8, y)
		for j := range a[i] {
			a[i][j] = uint8(i + j * i + j)
		}
	}
	return a
}

func main() {
	var a = SliceInit(3,2)
	fmt.Println(a)
} 
	


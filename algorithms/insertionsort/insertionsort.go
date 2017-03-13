package main

import "fmt"

func InsertionSort(a []int) {

	arrayLength := len(a)
	if arrayLength <= 1 {
		return
	}
	for i := 1; i < arrayLength; i++ {
		
		valueToInsert := a[i]
		insertPosition := i
	
		for insertPosition > 0 && a[insertPosition - 1] > valueToInsert {
		
			a[insertPosition] = a[insertPosition - 1]
			insertPosition--
		}
	
		a[insertPosition] = valueToInsert
	}
}

func main() {

        a := [][]int {{},{0},{99,56,12,90,1,4,11,67,23,45,13,10}, {4,2,8,1,5}, {89, 23, 12, 1, 2, 14, 56, 55, 55}}

        for _, x := range a {
                InsertionSort(x)
                fmt.Println(x)
        }
}

	  	
	
	

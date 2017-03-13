package main

import "fmt"

func MergeSort(a, temp []int, start, end int) {

	if start >= end {
		return
	}
	
	middle := (start + end) / 2
	MergeSort(a, temp, start, middle)
	MergeSort(a, temp, middle + 1, end)
	MergeHalves(a, temp, start, end)

}

func MergeHalves(a, temp []int, start, end int) {
	
	leftEnd := (start +  end) / 2
	rightStart := leftEnd + 1
	rightEnd := end
	size := end - start + 1

	left := start
	right := rightStart 
	index := start
	
	for left <= leftEnd && right <= rightEnd {
		if a[left] <= a[right] {
			temp[index] = a[left]
			left++
		} else {
			temp[index] = a[right]
			right++
		}
		index++
	}
	
	for left <= leftEnd {
		temp[index] = a[left]
		left++
		index++
	}
	for right <= rightEnd {
		temp[index] = a[right]
		right++
		index++
	}
	
	for size > 0 {
		a[start] = temp[start]
		start++
		size--
	}
}

func main() {
	
	a := [][]int {{99,56,12,90,1,4,11,67,23,45,13,10}, {4,2,8,1,5}, {89, 23, 12, 1, 2, 14, 56, 55, 55}}

        for _, x := range a {
		temp  := make([]int, len(x))
                MergeSort(x, temp, 0, len(x) - 1)
                fmt.Println(x)
        }
}	

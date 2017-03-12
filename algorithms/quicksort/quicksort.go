package main

import "fmt"

func QuickSort(a []int, left, right int) {


	if left >= right {
		return
	}

	pivot := a[(left + right) / 2]
	index := partition(a, pivot, left, right)
	QuickSort(a, left, index - 1)
	QuickSort(a, index, right)
 
}

func partition(a []int, pivot, left, right int) int {
	
	for left <= right {
	
		for a[left] < pivot {
			left++
		}

		for a[right] > pivot {
			right--
		}

		if left <= right {
			temp := a[left]
			a[left] = a[right]
			a[right] = temp
			left++
			right--
		}
	}
	return left 
}

func main() {
	
	var a = [][]int {{4,2,8,1,5}, {89, 23, 12, 1, 2, 14, 56, 55, 55}}

	for _, x := range a {
		QuickSort(x, 0, len(x) - 1)
		fmt.Println(x)
	}
}
			

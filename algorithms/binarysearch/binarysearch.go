package main

import "fmt"

func BinarySearch(a []int, x int) int {

	var lowerBound int = 0
	var upperBound int = len(a) - 1
        var midPoint int
        var index int


	for true {
		midPoint = lowerBound + (upperBound - lowerBound) / 2
		
		if a[midPoint] == x {
			index = midPoint
			break
		} else if a[midPoint] > x {
			upperBound = midPoint - 1
		} else if a[midPoint] < x {
			lowerBound = midPoint + 1
		}
		
		if upperBound < lowerBound {
			index = -1
                        break
		}
	}
	return index

}

func main() {
	
	var a = []int {1,4,6,7,9,12,14,16,30}
	var xValues = []int {0,1,9,16,26,30}
        for x:=0; x < len(xValues); x++ { 
		var index int = BinarySearch(a, xValues[x])
        	fmt.Printf("the index is: %d\n", index)
	}
}	

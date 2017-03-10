package main

import "fmt"
import "strings"

func WordCount(s string) int {

	var words = strings.Fields(s)
        var wordCount = make(map[string]int)
        for _, word := range words {
		wordCount[word] = wordCount[word] + 1
	} 
        fmt.Println(wordCount)
        return 0

}

func main() {

	s := "Test is test and test of Test test"	
	var wordCount = WordCount(s)
	fmt.Println(wordCount)

}


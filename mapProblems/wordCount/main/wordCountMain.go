package main

import (
	"fmt"
	"myGoMod/mapProblems/wordCount"
)

func main() {
	outputWordCount := wordCount.WordCount("When the going gets tough the tough gets going")
	fmt.Println("Count of words from a sentence is :", outputWordCount)
}

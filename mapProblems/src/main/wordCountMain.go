package main

import (
	"fmt"
	"myGoMod/mapProblems/src"
)

func main() {
	outputWordCount := src.WordCount("When the going gets tough the tough gets going")
	fmt.Println("Count of words from a sentence is ", outputWordCount)
}

package wordCount

import (
	"strings"
)

func WordCount(s string) map[string]int {
	var wordSlice []string = strings.Fields(s)
	wordMap := make(map[string]int)
	for _, word := range wordSlice {
		if wordCount, ok := wordMap[word]; ok {
			wordCount++
			wordMap[word] = wordCount
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
}

// Exercise: Maps
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

// You might find strings.Fields helpful.
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	words := strings.Fields(s)

	mapLetters := make(map[string]int)

	for _, word := range words {
		value, ok := mapLetters[word]
		if ok {
			mapLetters[word] = value + 1
		} else {
			mapLetters[word] = 1
		}
	}
	return mapLetters

}

func main() {
	wc.Test(WordCount)
}

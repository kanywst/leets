package main

import (
	"fmt"
	"math"
)

func main() {
	// wordsDict, word1, word2 := []string{"practice", "makes", "perfect", "coding", "makes"}, "makes", "coding"
	wordsDict, word1, word2 := []string{"a", "a", "b", "b"}, "a", "b"
	fmt.Println(shortestDistance(wordsDict, word1, word2))
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	w1, w2 := []int{}, []int{}
	for i, w := range wordsDict {
		if w == word1 {
			w1 = append(w1, i)

		}
		if w == word2 {
			w2 = append(w2, i)
		}
	}
	ans := math.MaxInt
	for _, ww1 := range w1 {
		for _, ww2 := range w2 {
			tmp := int(math.Abs(float64(ww1) - float64(ww2)))
			if tmp < ans {
				ans = tmp
			}
		}
	}
	return ans
}

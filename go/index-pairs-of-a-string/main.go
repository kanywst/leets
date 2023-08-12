package main

import (
	"fmt"
	"sort"
)

func main() {
	text, words := "thestoryofleetcodeandme", []string{"story", "fleet", "leetcode"}
	fmt.Println(indexPairs(text, words))
}

func indexPairs(text string, words []string) [][]int {
	var ans [][]int

	for _, w := range words {
		wLen := len(w)
		for i := 0; i+wLen <= len(text); i++ {
			if w == text[i:i+wLen] {
				ans = append(ans, []int{i, i + wLen - 1})
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i][0] == ans[j][0] {
			return ans[i][1] < ans[j][1]
		}
		return ans[i][0] < ans[j][0]
	})

	return ans
}

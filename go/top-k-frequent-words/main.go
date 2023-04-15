package main

import (
	"fmt"
	"sort"
)

func main() {
	words, k := []string{"i", "love", "leetcode", "i", "love", "coding"}, 2
	fmt.Println(topKFrequent(words, k))
}

func topKFrequent(words []string, k int) (ans []string) {
	m := make(map[string]int)
	for _, i := range words {
		if _, ok := m[i]; ok {
			m[i] += 1
		} else {
			m[i] = 1
		}
	}
	for key, val := range m {
		if val == k {
			ans = append(ans, key)
		}
	}
	sort.Strings(ans)
	return
}

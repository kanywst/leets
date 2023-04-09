package main

import "fmt"

func main() {
	words, s := []string{"a", "b", "c", "ab", "bc", "abc"}, "abc"
	fmt.Println(countPrefixes(words, s))
}

func countPrefixes(words []string, s string) (cnt int) {
	for _, w := range words {
		for i := 0; i < len(s); i++ {
			if s[:i+1] == w {
				cnt += 1
			}
		}
	}
	return
}

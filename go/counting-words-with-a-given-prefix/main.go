package main

import "fmt"

func main() {
	words, pref := []string{"pay", "attention", "practice", "attend"}, "at"
	fmt.Println(prefixCount(words, pref))
}

func prefixCount(words []string, pref string) (cnt int) {
	for _, w := range words {
		if len(w) >= len(pref) && w[:len(pref)] == pref {
			cnt += 1
		}
	}
	return
}

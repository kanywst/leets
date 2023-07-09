package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	keyboard, word := "abcdefghijklmnopqrstuvwxyz", "cba"
	fmt.Println(calculateTime(keyboard, word))
}

func calculateTime(keyboard string, word string) (ans int) {
	m := make(map[string]int)
	for i, j := range strings.Split(keyboard, "") {
		m[j] = i
	}
	ans, currentIndex := 0, 0
	for _, w := range strings.Split(word, "") {
		ans += int(math.Abs(float64(m[w]) - float64(currentIndex)))
		currentIndex = m[w]
	}
	return
}

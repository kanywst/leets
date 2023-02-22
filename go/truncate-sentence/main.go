package main

import (
	"fmt"
	"strings"
)

func main() {
	s, k := "Hello how are you Contestant", 4
	fmt.Println(truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	ss := strings.Split(s, " ")
	return strings.Join(ss[:k], " ")
}

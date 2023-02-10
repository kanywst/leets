package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 9
	fmt.Println(isStrictlyPalindromic(n))
}

func isStrictlyPalindromic(n int) bool {
	s, _ := strconv.ParseInt(strconv.Itoa(n), 2, 64)
	// for i := 0; i < len(s)/2; i++ {
	// 	if s[i:i+1] != s[len(s)-(i+1):len(s)-i] {
	// 		return false
	// 	}
	// }
	fmt.Println(s)
	return true
}

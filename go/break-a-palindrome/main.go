package main

import (
	"fmt"
)

func main() {
	// palindrome := "abccba"
	palindrome := "aa"
	fmt.Println(breakPalindrome(palindrome))
}

func breakPalindrome(palindromeStr string) string {
	if len(palindromeStr) <= 1 {
		return ""
	}

	palindrome := []byte(palindromeStr)
	for i := 0; i < len(palindrome)/2; i++ {
		if palindrome[i] == 'a' {
			continue
		}

		palindrome[i] = 'a'
		return string(palindrome)
	}

	palindrome[len(palindrome)-1] = 'b'
	return string(palindrome)
}

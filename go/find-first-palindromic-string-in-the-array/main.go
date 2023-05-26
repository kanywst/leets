package main

import "fmt"

func main() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(words))
}

func isPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

func firstPalindrome(words []string) (ans string) {
	for _, w := range words {
		if isPalindrome(w) {
			ans = w
			break
		}
	}
	return
}

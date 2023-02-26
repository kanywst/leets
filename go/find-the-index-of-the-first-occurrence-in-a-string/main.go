package main

import "fmt"

func main() {
	haystack, needle := "a", "a"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) (ans int) {
	n := len(needle)
	ans = -1
	for i := 0; i < len(haystack)-n+1; i++ {
		if haystack[i:i+n] == needle {
			ans = i
			break
		}
	}
	return
}

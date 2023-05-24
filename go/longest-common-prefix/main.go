package main

import "fmt"

func main() {
	strs := []string{"ab", "a"}
	fmt.Println("ans:", longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) (ans string) {
	if len(strs) == 1 {
		return strs[0]
	}
	minStr := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(minStr) {
			minStr = strs[i]
		}
	}
	for i := len(minStr); i > 0; i-- {
		f := true
		for _, s := range strs {
			fmt.Println(minStr[0:i], s[0:i])
			if minStr[0:i] != s[0:i] {
				f = false
				break
			}
		}
		if f {
			ans = minStr[0:i]
			break
		}
	}
	return
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence, searchWord := "i love eating burger", "burg"
	fmt.Println(isPrefixOfWord(sentence, searchWord))
}

func isPrefixOfWord(sentence string, searchWord string) (ans int) {
	ans = -1
	for i, j := range strings.Split(sentence, " ") {
		fmt.Println("Start", j)
		for k := 0; k < len(j); k++ {
			if len(j) > len(searchWord) {
				if j[0:len(searchWord)] == searchWord {
					fmt.Println(j[k : k+len(searchWord)])
					ans = i + 1
					break
				}
			} else {
				if searchWord == j {
					ans = i + 1
				}
			}
		}
		if ans != -1 {
			break
		}
	}
	return
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "!this  1-s b8d!"
	fmt.Println(countValidWords(sentence))
}

func countValidWords(sentence string) (cnt int) {
	for _, i := range strings.Split(sentence, " ") {
		tmp := false
		for _, j := range i {
			if j >= 97 && j <= 122 {
				tmp = true
			} else {
				tmp = false
				break
			}
		}
		if tmp {
			cnt += 1
		}
	}
	return
}

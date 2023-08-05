package main

import (
	"fmt"
	"strconv"
)

func main() {
	firstWord, secondWord, targetWord := "acb", "cba", "cdb"
	fmt.Println(isSumEqual(firstWord, secondWord, targetWord))
}

func translater(s string) int {
	tmp := ""
	for _, f := range s {
		tmp += strconv.Itoa(int(f - 97))
	}
	i, _ := strconv.Atoi(tmp)
	return i
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	if translater(firstWord)+translater(secondWord) == translater(targetWord) {
		return true
	}
	return false
}

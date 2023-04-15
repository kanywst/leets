package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := -1
	fmt.Println(toHex(num))
}

func toHex(num int) string {
	source := num

	hex := [16]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	answer := ""

	for source/16 != 0 {
		a := source % 16
		answer = hex[a] + answer
		source = source / 16
		if source < 16 {
			answer = strconv.Itoa(source) + answer
			break
		}
	}

	fmt.Println("answer:", answer)
	return ""
}

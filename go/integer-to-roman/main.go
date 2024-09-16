package main

import "fmt"

func main() {
	num := 3
	fmt.Println(intToRoman(num))
}

func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	return m[num]
}

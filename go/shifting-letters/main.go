package main

import (
	"fmt"
	"strings"
)

func main() {
	// s, shifts := "abc", []int{3, 5, 9}
	s, shifts := "bad", []int{10, 20, 30}
	fmt.Println(shiftingLetters(s, shifts))
}

func shiftingLetters(s string, shifts []int) string {
	tmp := strings.Split(s, "")
	for i := 0; i < len(tmp); i++ {
		ss := strings.Join(tmp[:i+1], "")
		fmt.Println(ss)
		for j, k := range ss {
			fmt.Println(i, j, k)
			tmp[j] = string(k + rune(shifts[i]))
		}
	}
	return strings.Join(tmp, "")
}

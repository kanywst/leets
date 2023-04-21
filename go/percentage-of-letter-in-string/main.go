package main

import "fmt"

func main() {
	s, letter := "foobar", "o"[0]
	fmt.Println(percentageLetter(s, letter))
}

func percentageLetter(s string, letter byte) int {
	cnt := 0
	for _, i := range s {
		if byte(i) == letter {
			cnt += 1
		}
	}
	return cnt * 100 / len(s)
}

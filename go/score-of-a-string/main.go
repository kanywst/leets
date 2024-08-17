package main

import (
	"fmt"
	"math"
)

func main() {
	s := "hello"
	fmt.Println(scoreOfString(s))
}
func scoreOfString(s string) int {
	cnt := 0
	for i := 0; i < len(s)-1; i++ {
		cnt += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}
	return cnt
}

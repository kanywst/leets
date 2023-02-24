package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A1:F1"
	fmt.Println(cellsInRange(s))
}

func cellsInRange(s string) (ans []string) {
	ss := strings.Split(s, ":")
	start, end := ss[0][0], ss[1][0]
	low, high := ss[0][1], ss[1][1]
	for i := start; i < end+1; i++ {
		for j := low; j < high+1; j++ {
			ans = append(ans, string(i)+string(j))
		}
	}
	return
}

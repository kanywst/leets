package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a, b := 1, 20
	fmt.Println(numberCount(a, b))
}

func numberCount(a int, b int) int {
	cnt := 0
	for i := a; i < b+1; i++ {
		s := strings.Split(strconv.Itoa(i), "")
		d := map[string]bool{}
		for _, v := range s {
			d[v] = true
		}
		if len(d) == len(s) {
			cnt += 1
		}
	}
	return cnt
}

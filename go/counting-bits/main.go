package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 2
	fmt.Println(countBits(n))
}

func countBits(n int) (ans []int) {
	for j := 0; j <= n; j++ {
		cnt := 0
		s := fmt.Sprintf("%b", j)
		for _, i := range strings.Split(s, "") {
			if "1" == i {
				cnt += 1
			}
		}
		ans = append(ans, cnt)
	}
	return
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 3
	fmt.Println(findNthDigit(n))
}

func findNthDigit(n int) (ans int) {
	tmp := []string{}
	for i := 1; i <= n; i++ {
		tmp = append(tmp, strconv.Itoa(i))
		if i == n {
			break
		}
	}
	tmp_ans := strings.Split(strings.Join(tmp, ""), "")
	ans, _ = strconv.Atoi(tmp_ans[n-1])
	return
}

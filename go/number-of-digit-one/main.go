package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 824883294
	fmt.Println(countDigitOne(n))
}

func countDigitOne(n int) (ans int) {
	tmp, tmpS := []int{}, []string{}
	for i := 0; i <= n; i++ {
		tmp = append(tmp, i)
	}
	for _, i := range tmp {
		tmpS = append(tmpS, strconv.Itoa(i))
	}
	for _, i := range strings.Join(tmpS, "") {
		if string(i) == "1" {
			ans += 1
		}
	}
	return
}

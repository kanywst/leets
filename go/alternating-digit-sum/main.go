package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 521
	fmt.Println(alternateDigitSum(n))
}

func alternateDigitSum(n int) (ans int) {
	tmp := strings.Split(strconv.Itoa(n), "")
	for i := 0; i < len(tmp); i++ {
		tmp_i, _ := strconv.Atoi(tmp[i])
		if i%2 == 0 {
			ans += tmp_i
		} else {
			ans -= tmp_i
		}
	}
	return
}

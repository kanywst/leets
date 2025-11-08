package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num, k := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}, 516
	fmt.Println(addToArrayForm(num, k))
}

func addToArrayForm(num []int, k int) (ans []int) {
	s := []string{}
	for _, i := range num {
		s = append(s, strconv.Itoa(i))
	}
	fmt.Println(s)
	n, _ := strconv.ParseUint(strings.Join(s, ""), 10, 64)
	kk, _ := strconv.ParseUint(strconv.Itoa(k), 10, 64)
	// for _, i := range n + kk {
	// 	fmt.Println(i)
	// tmp, _ := strconv.ParseUint(i, 10, 64)
	// ans = append(ans, unit64(tmp))
	// }
	return
}

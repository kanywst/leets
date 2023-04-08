package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	strs := []string{"alic3", "bob", "3", "4", "00000"}
	fmt.Println(maximumValue(strs))
}

func maximumValue(strs []string) int {
	ans := []int{}
	for _, s := range strs {
		n, e := strconv.Atoi(s)
		if e != nil {
			ans = append(ans, len(s))
		} else {
			ans = append(ans, n)
		}
	}
	sort.Ints(ans)
	return ans[len(ans)-1]
}

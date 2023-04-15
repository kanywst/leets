package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s, k := "leet2code3", 10
	fmt.Println(decodeAtIndex(s, k))
}

func decodeAtIndex(s string, k int) string {
	ind, ans := 0, ""
	for i, j := range strings.Split(s, "") {
		ss, err := strconv.Atoi(j)
		if err == nil {
			sss := s[ind:i]
			ind = i + 1
			ans += strings.Repeat(sss, ss)
			fmt.Println(j, sss, ss)
		}
	}
	return string(ans[k-1])
}

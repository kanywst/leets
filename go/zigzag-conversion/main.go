package main

import (
	"fmt"
	"strings"
)

func main() {
	s, numRows := "PAYPALISHIRING", 3
	fmt.Println(convert(s, numRows))
}

func convert(s string, numRows int) string {
	ss := strings.Split(s, "")
	ss_ans := make([]string, len(ss))
	for i := 0; i < len(ss); i++ {
		ss_ans[inumRows] = ss[i]
	}
	return strings.Join(ss_ans, "")
}

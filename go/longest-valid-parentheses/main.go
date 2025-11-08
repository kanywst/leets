package main

import (
	"fmt"
	"strings"
)

func main() {
	// s := "(()"
	// s := "()(())"
	s := "(()()"

	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) (cnt int) {
	ss := strings.Split(s, "")
	for {
		f := false
		for i := 0; i < len(ss)-1; i++ {
			// fmt.Println("ss:", ss[:i], ss[i+1:])
			// fmt.Println("strings.Join(ss[i:i+2]", strings.Join(ss[i:i+2], ""))
			// fmt.Println("cnt:", cnt)
			if strings.Join(ss[i:i+2], "") == "()" {
				fmt.Println("[before] if ss", ss, strings.Join(ss[i:i+2], ""))
				ss = append(ss[:i], ss[i+2:]...)
				fmt.Println("[after] if ss", ss, strings.Join(ss[i:i+2], ""))
				// if len(ss) == 2 && strings.Join(ss, "") == "()" {
				// 	fmt.Println("aaa")
				// 	f = false
				// 	break
				// }
				cnt += 1
				f = true
			}
		}
		if !f {
			break
		}
	}
	cnt *= 2
	return
}

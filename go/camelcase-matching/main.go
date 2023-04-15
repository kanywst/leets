package main

import "fmt"

func main() {
	queries, pattern := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"
	fmt.Println(camelMatch(queries, pattern))
}

func camelMatch(queries []string, pattern string) (ans []bool) {
	ans = make([]bool, len(queries))
	for i, q := range queries {
		tmp := ""
		for _, c := range q {
			if c >= 65 && c <= 90 {
				tmp += string(c)
			}
		}
		if tmp == pattern {
			ans[i] = true
		}
	}
	return
}

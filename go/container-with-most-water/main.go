package main

import (
	"fmt"
	"sort"
)

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func max_s(h []int) int {
	m := make(map[int]bool)
	uniq := []int{}

	for _, i := range h {
		if !m[i] {
			m[i] = true
			uniq = append(uniq, i)
		}
	}
	sort.Ints(uniq)
	return uniq[len(uniq)-2]
}

func maxArea(height []int) (ans int) {
	s, c, start := max_s(height), 0, 0
	fmt.Println("s:", s)
	for _, i := range height {
		fmt.Println("c,i:", c, i)
		if i >= s && start == 0 {
			start += 1
		}
		if i == s && c > 0 && start > 0 {
			break
		}
		if start > 0 {
			c += 1
		}
	}
	fmt.Println("s*c:", s, c)
	ans = s * c
	return
}

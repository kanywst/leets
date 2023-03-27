package main

import "fmt"

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	m := make(map[string]int)
	for _, i := range s {
		m[string(i)] += 1
	}
	for j, i := range s {
		if _, ok := m[string(i)]; ok {
			if m[string(i)] == 1 {
				return j
			}
		}
	}
	return -1
}

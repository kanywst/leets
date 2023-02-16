package main

import "fmt"

func main() {
	s, t := "egg", "add"
	// s, t := "badc", "baba"
	// s, t := "paper", "title"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		m1[int(s[i])] = i
		m2[int(t[i])] = i
	}
	return true
}

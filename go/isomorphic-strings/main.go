package main

import "fmt"

func main() {
	s, t := "egg", "add"
	// s, t := "badc", "baba"
	// s, t := "paper", "title"
	fmt.Println(isIsomorphic(s, t))
}

func isIsomorphic(s string, t string) bool {
	m := make(map[string]string)
	memo := make(map[string]string)
	for i := 0; i < len(s); i++ {
		if val, ok := m[string(s[i])]; ok {
			fmt.Println(m)
			if val != string(t[i]) {
				return false
			}
		} else if _, ok := memo[string(t[i])]; ok {
			return false
		} else {
			m[string(s[i])] = string(t[i])
			memo[string(t[i])] = ""
		}
	}
	return true
}

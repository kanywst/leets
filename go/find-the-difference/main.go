package main

import "fmt"

func main() {
	s, t := "abcd", "abcde"
	fmt.Println(findTheDifference(s, t))
}

func findTheDifference(s string, t string) byte {
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	for _, i := range s {
		m1[byte(i)] += 1
	}
	for _, i := range t {
		m2[byte(i)] += 1
	}
	for i, j := range m2 {
		if _, ok := m1[i]; ok {
			if j != m1[i] {
				return i
			}
		} else {
			return i
		}
	}
	return 0
}

package main

import "fmt"

func main() {
	s, part := "eemckxmckx", "emckx"
	fmt.Println("ans:", removeOccurrences(s, part))
}

func removeOccurrences(s string, part string) string {
	partLen, flag := len(part), false
	for {
		if len(s) < len(part) || flag || s == "" {
			break
		}
		for i := 0; i < len(s)-partLen+1; i++ {
			if s[i:i+partLen] == part {
				s = s[:i] + s[i+partLen:]
				break
			}
			if i == len(s)-partLen {
				flag = true
			}
		}
	}
	return s
}

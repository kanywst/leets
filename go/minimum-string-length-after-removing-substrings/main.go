package main

import "fmt"

func main() {
	s := "ABFCACDB"
	fmt.Println(minLength(s))
}

func minLength(s string) int {
	for {
		flag := false
		for i := 0; i < len(s)-1; i++ {
			if s[i:i+2] == "AB" || s[i:i+2] == "CD" {
				s = s[:i] + s[i+2:]
				flag = true
				break
			}
		}
		if !flag {
			break
		}
	}
	return len(s)
}

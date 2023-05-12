package main

import "fmt"

func main() {
	s := "s"
	fmt.Println("ans", makeGood(s))
}

func makeGood(s string) string {
	tmp := true
	for tmp {
		fmt.Println(s, "1")
		if len(s) == 0 || len(s) == 1 {
			break
		}
		for i := 0; i < len(s)-1; i++ {
			if s[i] >= 65 && s[i] <= 90 {
				if s[i] == s[i+1]-(97-65) {
					s = s[:i] + s[i+2:]
					break
				}
			}
			if s[i] >= 97 && s[i] <= 122 {
				if s[i] == s[i+1]+(97-65) {
					s = s[:i] + s[i+2:]
					break
				}
			}
			if i == len(s)-2 {
				tmp = false
			}
		}
	}
	return s
}

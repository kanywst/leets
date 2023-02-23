package main

import (
	"fmt"
	"strings"
)

func main() {
	title := "First leTTeR of EACH Word"
	fmt.Println(capitalizeTitle(title))
}

func capitalizeTitle(title string) string {
	s := strings.Title(strings.ToLower(title))
	tmp_s := strings.Split(s, " ")
	for i, j := range tmp_s {
		if len(j) < 3 {
			tmp_s[i] = strings.ToLower(j)
		}
	}
	return strings.Join(tmp_s, " ")
}

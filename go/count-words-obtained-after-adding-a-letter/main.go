package main

import "fmt"

func main() {
	startWords, targetWords := []string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}
	fmt.Println(wordCount(startWords, targetWords))
}

func wordCount(startWords []string, targetWords []string) (cnt int) {
	m := make(map[string]bool)
	for i, _ := range startWords {
		for _, j := range targetWords {
			if len(startWords[i]) > len(j) {
				continue
			}
			if startWords[i] == j[:len(startWords[i])] {
				if _, ok := m[startWords[i]]; !ok {
					cnt += 1
				}
				m[startWords[i]] = true
				fmt.Println(startWords[i], j[:len(startWords[i])])
			}
		}
	}
	return
}

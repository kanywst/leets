package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	dictionary, sentence := []string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
}

func replaceWords(dictionary []string, sentence string) string {
	tmp_s := strings.Split(sentence, " ")
	sort.Strings(dictionary)
	for i, j := range tmp_s {
		for _, s := range dictionary {
			if len(j) >= len(s) && s == j[0:len(s)] {
				tmp_s[i] = s
				break
			}
		}
	}
	return strings.Join(tmp_s, " ")
}

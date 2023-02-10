package main

import "fmt"

func main() {
	sentence := "thequickbrownfoxjumpsoverthelazydog"
	fmt.Println(checkIfPangram(sentence))
}

func checkIfPangram(sentence string) bool {
	m := make(map[string]string)
	for _, y := range sentence {
		if _, ok := m[string(y)]; !ok {
			m[string(y)] = "nil"
		}
	}
	if len(m) == 26 {
		return true
	} else {
		return false
	}
}

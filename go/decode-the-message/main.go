package main

import "fmt"

func main() {
	//key1, message1 := "eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
	key2, message2 := "the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"
	fmt.Println(decodeMessage(key2, message2))
}

func decodeMessage(key string, message string) string {
	m := make(map[string]string)
	c := 97
	for _, j := range key {
		if _, ok := m[string(j)]; !ok {
			if string(j) != " " {
				m[string(j)] = string(c)
				c += 1
			}
		}
	}
	answer := ""
	for _, j := range message {
		if string(j) == " " {
			answer += " "
		} else {
			answer += m[string(j)]
		}
	}
	return answer
}

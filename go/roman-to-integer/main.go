package main

import "fmt"

func main() {
	// s := "LVIII"
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) (ans int) {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for _, v := range s {
		ans += m[string(v)]
	}
	return
}

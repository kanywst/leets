package main

import "fmt"

func main() {
	s := "([{}])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	sum := 0
	for _, i := range s {
		if i == 41 && sum != 0 {
			sum += -40
		} else if i == 93 && sum != 0 {
			sum += -91
		} else if i == 125 && sum != 0 {
			sum += -123
		} else {
			sum += int(i)
		}
	}
	if sum == 0 {
		return true
	} else {
		return false
	}
}

package main

import "fmt"

func main() {
	maxChoosableInteger, desiredTotal := 10, 11
	fmt.Println(canIWin(maxChoosableInteger, desiredTotal))
}

func canIWin(maxChoosableInteger int, desiredTotal int) (ans bool) {
	fmt.Println(6 % 4)
	if desiredTotal%maxChoosableInteger == 0 {
		ans = false
	} else {
		ans = true
	}
	return
}

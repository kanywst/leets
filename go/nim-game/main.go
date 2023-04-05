package main

import "fmt"

func main() {
	n := 4
	fmt.Println(canWinNim(n))
}

func canWinNim(n int) (ans bool) {
	return n%4 != 0
}

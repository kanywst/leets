package main

import "fmt"

func main() {
	columnNumber := 701
	fmt.Println(convertToTitle(columnNumber))
}

func convertToTitle(columnNumber int) (res string) {
	for columnNumber > 0 {
		columnNumber--
		res = string(byte(columnNumber%26)+'A') + res
		columnNumber /= 26
	}
	return
}

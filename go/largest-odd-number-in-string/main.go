package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := "32782489638346578713315098393010310518347382"
	fmt.Println(largestOddNumber(num))
}

func largestOddNumber(num string) (ans string) {
	// n, _ := strconv.Atoi(num)
	for i := len(num); i >= 0; i-- {
		num = num[:i]
		n, _ := strconv.Atoi(num)
		if n%2 == 1 {
			fmt.Println(num)
			ans = num
			break
		}
	}
	return
}

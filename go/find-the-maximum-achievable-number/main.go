package main

import "fmt"

func main() {
	num, t := 4, 1
	fmt.Println(theMaximumAchievableX(num, t))
}

func theMaximumAchievableX(num int, t int) int {
	return num + (t * 2)
}

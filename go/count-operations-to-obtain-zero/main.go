package main

import "fmt"

func main() {
	num1, num2 := 2, 3
	fmt.Println(countOperations(num1, num2))
}

func countOperations(num1 int, num2 int) (cnt int) {
	for num1 > 0 && num2 > 0 {
		cnt += 1
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
	return
}

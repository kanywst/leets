package main

import "fmt"

func main() {
	k := 7
	fmt.Println(findMinFibonacciNumbers(k))
}

func fibonacci() func() int {
	f0, f1 := 0, 1

	return func() (x int) {
		x = f0
		f0, f1 = f1, f0+f1
		return
	}
}

func findMinFibonacciNumbers(k int) int {
	f := fibonacci()
	for i := 0; i < k; i++ {
		fmt.Println(f())
	}
	return 0
}

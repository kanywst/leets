package main

import "fmt"

func main() {
	n := 0
	fmt.Println(isUgly(n))
}

func isUgly(n int) bool {
	ans := false
	ugly_num := []int{2, 3, 5}
	if n == 1 {
		return true
	}
	if n == 0 {
		return false
	}
	for _, i := range ugly_num {
		if n%i == 0 {
			ans = isUgly(n / i)
			break
		}
	}
	return ans
}

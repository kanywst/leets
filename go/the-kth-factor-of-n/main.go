package main

import "fmt"

func main() {
	n, k := 12, 3
	fmt.Println(kthFactor(n, k))
}

func factor(n int) (ans []int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ans = append(ans, i)
		}
	}
	return
}
func kthFactor(n int, k int) int {
	nn := factor(n)
	if len(nn) <= k-1 {
		return -1
	} else {
		return nn[k-1]
	}
}

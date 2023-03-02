package main

import (
	"fmt"
)

func main() {
	// n := 8 // 6
	n := 1
	fmt.Println(pivotInteger(n))
}

func sum(slice []int) (sum int) {
	for _, x := range slice {
		sum += x
	}
	return
}

func arr(start int, end int) (ans []int) {
	for i := 0; i <= end-start+1; i++ {
		ans = append(ans, start+i)
	}
	return
}

func pivotInteger(n int) (ans int) {
	if n == 1 {
		ans = 1
		return
	}
	for i := 1; i < n; i++ {
		if sum(arr(1, i-1)) == sum(arr(i, n-1)) {
			ans = i
			return
		}
	}
	ans = -1
	return
}

package main

import "fmt"

func main() {
	n := 11
	fmt.Println(nthUglyNumber(n))
}

func nthUglyNumber(n int) int {
	cnt := 1
	ans := []int{}
	for i := 1; cnt <= n+10; i++ {
		if cnt == 1 || i%2 == 0 || i%3 == 0 || i%5 == 0 || (i%2 == 0 && i%3 == 0) || (i%2 == 0 || i%5 == 0) || (i%3 == 0 && i%5 == 0) {
			cnt = i
			ans = append(ans, cnt)
		}
	}
	fmt.Println(ans)
	return ans[n-1]
}

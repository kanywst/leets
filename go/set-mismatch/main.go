package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(nums))
}

func findErrorNums(nums []int) (ans []int) {
	m := make(map[int]int)
	for _, n := range nums {
		m[n] += 1
	}
	for k, v := range m {
		if v >= 2 {
			ans = append(ans, k)
		}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			ans = append(ans, i)
		}
	}
	return
}

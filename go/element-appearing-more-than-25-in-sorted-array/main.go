package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 3}
	// arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	fmt.Println(findSpecialInteger(arr))
}

func findSpecialInteger(arr []int) int {
	m := make(map[int]int)
	for _, i := range arr {
		m[i] += 1
	}
	sub := []int{}
	for k, v := range m {
		if v*100/len(arr) > 25 {
			return k
		}
		if v*100/len(arr) == 25 {
			sub = append(sub, k)
		}
	}
	return sub[0]
}

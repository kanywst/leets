package main

import "fmt"

func main() {
	// arr, k := []int{2, 3, 4, 7, 11}, 5
	arr, k := []int{1, 2, 3, 4}, 2
	fmt.Println(findKthPositive(arr, k))
}

func findKthPositive(arr []int, k int) int {
	ans := []int{}
	cnt, i := 1, 0
	for {
		if i >= len(arr) {
			break
		}
		if arr[i] == cnt {
			i++
			cnt++
		} else if arr[i] > cnt {
			ans = append(ans, cnt)
			cnt++
		}
	}
	for c := cnt; k > len(ans); c++ {
		fmt.Println("c")
		ans = append(ans, c)
	}
	return ans[k-1]
}

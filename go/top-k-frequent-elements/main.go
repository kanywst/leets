package main

import "fmt"

func main() {
	// nums, k := []int{1, 1, 1, 2, 2, 3}, 2
	nums, k := []int{-1, -1}, 1
	fmt.Println(topKFrequent(nums, k))
}

func topKFrequent(nums []int, k int) (ans []int) {
	cnt_map := make(map[int]int)
	cnt_slice := make([][]int, len(nums)+1)
	tmp := nums[0]
	for _, j := range nums {
		if j != tmp {
			tmp = j
		}
		if j == tmp {
			if _, ok := cnt_map[tmp]; !ok {
				cnt_map[tmp] = 0
			}
			cnt_map[tmp] += 1
		}
	}
	for i, j := range cnt_map {
		cnt_slice[j] = append(cnt_slice[j], i)
	}
	for i := len(cnt_slice) - 1; i > 0; i-- {
		ans = append(ans, cnt_slice[i]...)
		if len(ans) == k {
			break
		}
	}
	return
}

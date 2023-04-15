package main

import (
	"fmt"
	"sort"
)

func main() {
	// arr := []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
	// arr := []int{7}
	arr := []int{6, 1, 9}
	fmt.Println(minJumps(arr))
}

func minJumps(arr []int) (cnt int) {
	start, end, cnt := arr[0], arr[len(arr)-1], 0
	if len(arr) == 1 {
		return
	}
	m := make(map[int]int)
	for _, i := range arr {
		if _, ok := m[i]; ok {
			m[i] += 1
		} else {
			m[i] = 0
		}
	}
	sort.Ints(arr)
	for _, i := range arr {
		fmt.Println(i, start, end, cnt)
		if i == end {
			if m[i] == 0 {
				break
			} else {
				m[i] -= 1
				cnt += 1
			}
		}
		if i >= start {
			cnt += 1
		}
	}
	cnt--
	return
}

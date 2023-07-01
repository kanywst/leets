package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1, arr2, arr3 := []int{1, 2, 3, 4, 5}, []int{1, 2, 5, 7, 9}, []int{1, 3, 4, 5, 8}
	fmt.Println(arraysIntersection(arr1, arr2, arr3))
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) (ans []int) {
	m := make(map[int]int)
	arr := []int{}
	arr = append(arr, arr1...)
	arr = append(arr, arr2...)
	arr = append(arr, arr3...)
	for _, a := range arr {
		m[a] += 1
	}
	for k, v := range m {
		if v == 3 {
			ans = append(ans, k)
		}
	}
	sort.Ints(ans)
	return
}

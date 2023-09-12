package main

import "fmt"

func main() {
	nums := [][]int{[]int{3, 6}, []int{1, 5}, []int{4, 7}}
	fmt.Println(numberOfPoints(nums))
}

func uniq(arr []int) (uniq []int) {
	m := make(map[int]bool)

	for _, ele := range arr {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return
}

func numberOfPoints(nums [][]int) int {
	res := []int{}
	for _, ns := range nums {
		for i := ns[0]; i <= ns[1]; i++ {
			res = append(res, i)
		}
	}
	return len(uniq(res))
}

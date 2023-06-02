package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	nums := []int{1, 2, 10, 5, 7}
	fmt.Println(canBeIncreasing(nums))
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

func canBeIncreasing(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		s := []int{}
		s = append(s, nums[:i]...)
		s = append(s, nums[i+1:]...)
		t := make([]int, len(s))
		copy(t, s)
		sort.Ints(t)
		if reflect.DeepEqual(s, t) && len(uniq(s)) == len(s) {
			return true
		}
	}
	return false
}

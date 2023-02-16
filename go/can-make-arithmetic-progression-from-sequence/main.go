package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 4}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func canMakeArithmeticProgression(arr []int) (ans bool) {
	sort.Ints(arr)
	ans = true
	if len(arr) < 2 {
		ans = false
	}
	tmp := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if tmp != arr[i+1]-arr[i] {
			ans = false
			break
		}
		tmp = arr[i+1] - arr[i]
	}
	return
}

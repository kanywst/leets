package main

import "fmt"

func main() {
	arr1, arr2 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}, []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) (ans []int) {
	// m:= make(map[int]bool), false, []int{}
	for _, j := range arr2 {
		n := 0
		for _, i := range arr1 {
			fmt.Println(i, j, n, arr1)
			if i == j {
				ans = append(ans, i)
				arr1 = append(arr1[:n], arr1[n+1:]...)
			} else {
				n++
			}
		}
	}
	// ans = append(ans, arr1...)
	return
}

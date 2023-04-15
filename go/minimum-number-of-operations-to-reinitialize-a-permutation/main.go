package main

import "fmt"

func main() {
	n := 4
	fmt.Println(reinitializePermutation(n))
}

func reinitializePermutation(n int) int {
	perm := []int{}
	for i := 0; i < n; i++ {
		perm = append(perm, i)
	}
	arr := make([]int, n)
	copy(arr, perm)
	fmt.Println(arr, perm)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			t1, t2 := perm[i], perm[i/2]
			arr[i/2] = t1
			arr[i] = t2
		}
		if i%2 == 1 {
			t1, t2 := perm[i], perm[n/2+(i-1)/2]
			arr[n/2+(i-1)/2] = t1
			arr[i] = t2
		}
		fmt.Println(i, arr)
	}
	return 0
}

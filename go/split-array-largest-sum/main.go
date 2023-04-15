package main

import "fmt"

func main() {
	// nums, k := []int{7, 2, 5, 10, 8}, 2
	nums, k := []int{1, 4, 4}, 3
	fmt.Println(splitArray(nums, k))
}

func sum(x []int) (y int) {
	for _, i := range x {
		y += i
	}
	return
}

func splitArray(nums []int, k int) (ans int) {
	if len(nums) >= k {
		ans = sum(nums)
		return
	}
	fmt.Println(k)
	fmt.Println(nums[0:k+1], nums[k+1:])
	n1, n2 := sum(nums[0:k+1]), sum(nums[k+1:])
	if n1 > n2 {
		ans = n1
	} else {
		ans = n2
	}
	return
}

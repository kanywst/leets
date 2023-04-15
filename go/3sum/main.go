package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) (ans [][]int) {
	m := make(map[int]bool)
	uniq := []int{}

	for _, ele := range nums {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	for i := 0; i < len(uniq); i++ {
		for j := 0; j < len(uniq); j++ {
			for k := 0; k < len(uniq); k++ {
				if i != j && j != k && i != k {
					fmt.Println(uniq[i], uniq[j], uniq[k])
					if uniq[i]+uniq[j]+uniq[k] == 0 {
						ans = append(ans, []int{uniq[i], uniq[j], uniq[k]})
					}
				}
			}
		}
	}
	return
}

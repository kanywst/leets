package main

import "fmt"

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}

func findMaxConsecutiveOnes(nums []int) int {
	c := false
	tmp := 0
	ans := 0
	for _, n := range nums {
		if n == 1 && c {
			tmp += 1
		}
		if n == 1 && !c {
			tmp += 1
			c = true
		}
		if n != 1 && c {
			if ans < tmp {
				ans = tmp
			}
			c = false
			tmp = 0
		}
	}
	if c {
		if ans < tmp {
			ans = tmp
		}
	}
	return ans
}

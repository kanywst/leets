package main

import "fmt"

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Println(i)
			tmp_ans := 0
			for n := range arr[:i] {
				tmp_ans += n
			}
			ans += tmp_ans
		}
	}
	return ans
}

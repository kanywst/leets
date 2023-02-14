package main

import "fmt"

func main() {
	numbers, target := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) (ans []int) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				ans = []int{i + 1, j + 1}
				break
			}
		}
	}
	return
}

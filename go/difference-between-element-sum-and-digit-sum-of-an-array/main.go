package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func differenceOfSum(nums []int) int {
	sum := 0
	digit_sum := 0
	new_nums := []string{}

	for _, x := range nums {
		sum += x

		xx := strings.Split(strconv.Itoa(x), "")
		new_nums = append(new_nums, xx...)
	}

	for _, y := range new_nums {
		i, _ := strconv.Atoi(y)
		digit_sum += i
	}

	return int(math.Abs(float64(sum - digit_sum)))
}

func main() {
	nums := []int{1, 15, 6, 3}
	fmt.Println(differenceOfSum(nums))
}

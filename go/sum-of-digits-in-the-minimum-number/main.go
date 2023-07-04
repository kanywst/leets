package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []int{34, 23, 1, 24, 75, 33, 54, 8}
	fmt.Println(sumOfDigits(nums))
}

func sumOfDigits(nums []int) int {
	sort.Ints(nums)
	s := strings.Split(strconv.Itoa(nums[0]), "")
	tmp := 0
	for _, i := range s {
		t, _ := strconv.Atoi(string(i))
		tmp += t
	}
	if tmp%2 == 0 {
		return 1
	}
	return 0
}

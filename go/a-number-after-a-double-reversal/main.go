package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	num := 526
	fmt.Println(isSameAfterReversals(num))
}

func isSameAfterReversals(num int) (ans bool) {
	nums := strings.Split(strconv.Itoa(num), "")
	tmp := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		tmp[len(nums)-i-1] = nums[i]
	}
	nums_t, _ := strconv.Atoi(strings.Join(tmp, ""))
	tmp_t, _ := strconv.Atoi(strings.Join(nums, ""))
	ans = false
	if len(strings.Split(strconv.Itoa(nums_t), "")) == len(strings.Split(strconv.Itoa(tmp_t), "")) {
		ans = true
	}
	return
}

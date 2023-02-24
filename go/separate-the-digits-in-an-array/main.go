package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{13, 25, 83, 77}
	fmt.Println(separateDigits(nums))
}

func separateDigits(nums []int) (ans []int) {
	tmp := []string{}
	for _, i := range nums {
		tmp = append(tmp, strings.Split(strconv.Itoa(i), "")...)
	}
	for _, i := range tmp {
		tmp_i, _ := strconv.Atoi(i)
		ans = append(ans, tmp_i)
	}
	return ans
}

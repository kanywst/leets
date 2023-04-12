package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := "014455"
	fmt.Println(largestGoodInteger(num))
}

func largestGoodInteger(num string) string {
	tmp, cnt, n := num[0], 0, []int{}
	for i := 1; i < len(num); i++ {
		if tmp == num[i] {
			cnt += 1
		} else {
			cnt = 0
		}
		if cnt == 2 {
			t := string(tmp) + string(tmp) + string(tmp)
			tt, _ := strconv.Atoi(t)
			n = append(n, tt)
			cnt = 0
		}
		tmp = num[i]
	}
	if len(n) == 0 {
		return ""
	}
	sort.Ints(n)
	ans := strconv.Itoa(n[len(n)-1])
	if ans == "0" {
		return "000"
	}
	return ans
}

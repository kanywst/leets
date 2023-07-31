package main

import (
	"fmt"
	"sort"
)

func main() {
	hours, target := []int{0, 1, 2, 3, 4}, 2
	fmt.Println(numberOfEmployeesWhoMetTarget(hours, target))
}

func numberOfEmployeesWhoMetTarget(hours []int, target int) (ans int) {
	sort.Ints(hours)
	for _, h := range hours {
		if h >= target {
			ans += 1
		}
	}
	return
}

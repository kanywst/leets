package main

import (
	"fmt"
	"sort"
)

func main() {
	weight := []int{900, 950, 800, 1000, 700, 800}
	fmt.Println(maxNumberOfApples(weight))
}

func maxNumberOfApples(weight []int) (cnt int) {
	sort.Ints(weight)
	tmp := 0
	for _, w := range weight {
		tmp += w
		if tmp > 5000 {
			return
		}
		cnt += 1
	}
	return
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))
}

func average(salary []int) float64 {
	sort.Slice(salary, func(i, j int) bool {
		return salary[i] < salary[j]
	})
	ss := salary[1 : len(salary)-1]
	sum := 0
	for i := 0; i < len(ss); i++ {
		sum += ss[i]
	}
	return float64(sum) / float64(len(ss))
}

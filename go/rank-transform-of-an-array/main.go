package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(arr))
}

func arrayRankTransform(arr []int) (ans []int) {
	tempArr := []int{}
	m := make(map[int]int)
	for _, a := range arr {
		tempArr = append(tempArr, a)
	}
	sort.Ints(tempArr)
	i := 0
	for _, j := range tempArr {
		if _, ok := m[j]; !ok {
			m[j] = i + 1
			i += 1
		}
	}
	for _, i := range arr {
		ans = append(ans, m[i])
	}
	return
}

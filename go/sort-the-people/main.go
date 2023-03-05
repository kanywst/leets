package main

import (
	"fmt"
	"sort"
)

func main() {
	names, heights := []string{"Mary", "John", "Emma"}, []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))
}

func sortPeople(names []string, heights []int) (ans []string) {
	m := make(map[int]string)
	// ans := make([]string, len(names))
	for i, j := range names {
		m[heights[i]] = j
	}
	sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	for _, i := range heights {
		ans = append(ans, m[i])
	}
	return
}

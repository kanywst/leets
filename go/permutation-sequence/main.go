package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n, k := 3, 3
	fmt.Println(getPermutation(n, k))
}

func getPermutation(n int, k int) string {
	x := []int{}
	for i := 1; i <= n; i++ {
		x = append(x, i)
	}
	perm, tmp := permutation(x), []string{}
	fmt.Println(perm)
	for _, i := range perm[k-1] {
		tmp = append(tmp, strconv.Itoa(i))
	}
	return strings.Join(tmp, "")
}

func permutation(vids [][]int) [][]int {
	toRet := [][]int{}

	if len(vids) == 0 {
		return toRet
	}

	if len(vids) == 1 {
		for _, vid := range vids[0] {
			toRet = append(toRet, []int{vid})
		}
		return toRet
	}

	t := permutation(vids[1:])
	for _, vid := range vids[0] {
		for _, perm := range t {
			toRetAdd := append([]int{vid}, perm...)
			toRet = append(toRet, toRetAdd)
		}
	}

	return toRet
}

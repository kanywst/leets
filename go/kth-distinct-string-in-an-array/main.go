package main

import "fmt"

func main() {
	// arr, k := []string{"d", "b", "c", "b", "c", "a"}, 2
	arr, k := []string{"aaa", "aa", "a"}, 1
	fmt.Println(kthDistinct(arr, k))
}

func kthDistinct(arr []string, k int) (r string) {
	m := make(map[string]int)
	ans := []string{}
	for _, i := range arr {
		m[i] += 1
	}
	for _, i := range arr {
		if m[i] == 1 {
			ans = append(ans, i)
		}
	}
	if len(ans) < k {
		r = ""
	} else {
		r = ans[k-1]
	}
	return
}

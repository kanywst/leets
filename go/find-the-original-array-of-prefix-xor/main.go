package main

import "fmt"

func main() {
	pref := []int{5, 2, 0, 3, 1}
	fmt.Println(findArray(pref))
}

func findArray(pref []int) []int {
	tmp := pref[0]
	ans := make([]int, len(pref))
	for i := 1; i < len(pref); i++ {
		fmt.Println(tmp, pref[i])
		tmp = tmp ^ pref[i]
		ans = append(ans, tmp)
	}
	return ans
}

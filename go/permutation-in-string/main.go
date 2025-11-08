package main

import (
	"fmt"
	"strings"
)

func main() {
	s1, s2 := "a", "ab"
	fmt.Println(checkInclusion(s1, s2))
}

func remove(ar []int, i int) []int {
	tmp := make([]int, len(ar))
	copy(tmp, ar)
	return append(tmp[0:i], tmp[i+1:]...)
}

func permutation_full(ar []int) [][]int {
	var result [][]int
	if len(ar) == 1 {
		return append(result, ar)
	}
	for i, a := range ar {
		for _, b := range permutation_full(remove(ar, i)) {
			result = append(result, append([]int{a}, b...))
		}
	}
	return result
}

func combination(ar []int, n int) (result [][]int) {
	if n <= 0 || len(ar) < n {
		return
	}
	if n == 1 {
		for _, a := range ar {
			result = append(result, []int{a})
		}
	} else if len(ar) == n {
		result = append(result, ar)
	} else {
		for _, a := range combination(ar[1:], n-1) {
			result = append(result, append([]int{ar[0]}, a...))
		}
		result = append(result, combination(ar[1:], n)...)
	}
	return
}

func permutation(ar []int, n int) (result [][]int) {
	for _, a := range combination(ar, n) {
		result = append(result, permutation_full(a)...)
	}
	return
}

func checkInclusion(s1 string, s2 string) (ans bool) {
	s1_l, s2_l, ss1 := strings.Split(s1, ""), strings.Split(s2, ""), []string{}
	for i := 0; i < len(s1_l); i++ {
		ss1 = append(ss1, s1_l[len(s1_l)-i-1])
	}
	s := permutation(s1_l, len(s1_l))
	fmt.Println("len(s2_l)-len(s1_l)-1:", len(s2_l)-len(s1_l)-1)
	for i := 0; i < len(s2_l)-len(s1_l); i++ {
		fmt.Println(i)
		fmt.Println(strings.Join(s2_l[i:i+len(s1_l)], ""), "==", s1, "||", strings.Join(s2_l[i:i+len(s1_l)], ""), "==", strings.Join(ss1, ""))
		if strings.Join(s2_l[i:i+len(s1_l)], "") == s1 || strings.Join(s2_l[i:i+len(s1_l)], "") == strings.Join(ss1, "") {
			ans = true
			break
		}
	}
	return
}

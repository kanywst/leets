package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(countElements(arr))
}

func countElements(arr []int) (ans int) {
	m := make(map[int]bool)
	for _, a := range arr {
		m[a] = true
	}
	for _, a := range arr {
		if _, ok := m[a+1]; ok {
			ans += 1
		}
	}
	return
}

package main

import "fmt"

func main() {
	arr := []int{10, 2, 5, 3}
	fmt.Println(checkIfExist(arr))
}

func checkIfExist(arr []int) (ans bool) {
	tmp_ans := []int{}
	for _, i := range arr {
		tmp_ans = append(tmp_ans, i*2)
	}
	for i := 0; i < len(tmp_ans); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && tmp_ans[i] == arr[j] {
				ans = true
				break
			}
		}
		if ans {
			break
		}
	}
	return
}

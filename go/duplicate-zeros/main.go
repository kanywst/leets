package main

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
}

func duplicateZeros(arr []int) {
	ans := []int{}
	for _, a := range arr {
		if a == 0 {
			ans = append(ans, 0)
		}
		ans = append(ans, a)
	}
	ans = ans[:len(arr)]
	for i := 0; i < len(arr); i++ {
		arr[i] = ans[i]
	}
}

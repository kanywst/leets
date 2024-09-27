package main

import (
	"fmt"
	"sort"
	"strconv"
)

// https://spacedleet.vercel.app/solutions/remove-k-digits/go

func main() {
	num, k := "10001", 4
	fmt.Println(removeKdigits(num, k))
}

func removeKdigits(num string, k int) (ans string) {
	tmp := []int{}
	for i := 0; i < len(num)-k+1; i++ {
		fmt.Println(num[:i], num[i+k:])
		n, _ := strconv.Atoi(num[:i] + num[i+k:])
		tmp = append(tmp, n)
	}
	sort.Ints(tmp)
	fmt.Println(tmp)
	if len(tmp) == 0 {
		ans = "0"
	} else {
		ans = strconv.Itoa(tmp[0])
	}
	return
}

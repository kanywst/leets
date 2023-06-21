package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	num := 4325
	fmt.Println(splitNum(num))
}

func splitNum(num int) int {
	s := strings.Split(strconv.Itoa(num), "")
	t := make([]int, len(s))
	for i, j := range s {
		t[i], _ = strconv.Atoi(j)
	}
	sort.Ints(t)
	t1, t2 := "", ""
	for i := 0; i < len(t)/2; i++ {
		t1 += strconv.Itoa(t[i*2])
		t2 += strconv.Itoa(t[i*2+1])
	}
	if len(s)%2 == 1 {
		t1 += strconv.Itoa(t[len(s)-1])
	}
	tt1, _ := strconv.Atoi(t1)
	tt2, _ := strconv.Atoi(t2)
	return tt1 + tt2
}

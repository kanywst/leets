package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	num := 2932
	fmt.Println(minimumSum(num))
}

func minimumSum(num int) int {
	s := strings.Split(strconv.Itoa(num), "")
	sort.Strings(s)
	tmp_1, _ := strconv.Atoi(s[0] + s[2])
	tmp_2, _ := strconv.Atoi(s[1] + s[3])
	return tmp_1 + tmp_2
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 234
	fmt.Println(subtractProductAndSum(n))

}

func subtractProductAndSum(n int) int {
	p, s := 1, 0
	for _, i := range strings.Split(strconv.Itoa(n), "") {
		r, _ := strconv.Atoi(i)
		p *= r
		s += r
	}
	return p - s
}

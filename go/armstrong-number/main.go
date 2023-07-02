package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	n := 153
	fmt.Println(isArmstrong(n))
}

func isArmstrong(n int) bool {
	tmp := 0
	s := strings.Split(strconv.Itoa(n), "")
	for _, i := range s {
		t, _ := strconv.Atoi(i)
		tmp += int(math.Pow(float64(t), float64(len(s))))
	}
	return tmp == n
}

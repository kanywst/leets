package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, b := "1010", "1011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	aa, _ := strconv.ParseInt(a, 2, 32)
	bb, _ := strconv.ParseInt(b, 2, 32)
	fmt.Println(aa + bb)
	return ""
}

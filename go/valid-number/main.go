package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := ".1"
	fmt.Println(isNumber(s))
}
func isNumber(s string) bool {

	_, e := strconv.Atoi(strings.Trim(s, "."))
	if e != nil {
		return false
	}
	return true
}

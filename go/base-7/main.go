package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 100
	fmt.Println(convertToBase7(num))
}

func convertToBase7(num int) string {
	return fmt.Sprintf(strconv.FormatInt(int64(num), 7))
}

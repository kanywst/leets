package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num1, num2 := "2", "3"
	fmt.Println(multiply(num1, num2))
}

func multiply(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	big_n1 := big.NewInt(int64(n1))
	big_n2 := big.NewInt(int64(n2))
	ans := big_n1.Div(big_n1, big_n2)
	return ans.String()
}

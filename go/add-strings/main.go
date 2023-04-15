package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num1, num2 := "3876620623801494171", "6529364523802684779"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)
	big_n1 := big.NewInt(int64(n1))
	big_n2 := big.NewInt(int64(n2))
	sum := big_n1.Add(big_n1, big_n2)
	return sum.String()
}

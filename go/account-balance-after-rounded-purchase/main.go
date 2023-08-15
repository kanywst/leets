package main

import (
	"fmt"
)

func main() {
	purchaseAmount := 9
	fmt.Println(accountBalanceAfterPurchase(purchaseAmount))
}

func roundToNearest(num int) int {
	if num%10 >= 5 {
		return ((num / 10) + 1) * 10
	}
	return (num / 10) * 10
}

func accountBalanceAfterPurchase(purchaseAmount int) int {
	return 100 - roundToNearest(purchaseAmount)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	prices, money := []int{1, 2, 2}, 3
	fmt.Println(buyChoco(prices, money))
}

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	p := prices[0] + prices[1]
	if p > money {
		return money
	}
	return money - p
}

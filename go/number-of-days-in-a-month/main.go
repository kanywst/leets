package main

import "fmt"

func main() {
	year, month := 1992, 7
	fmt.Println(numberOfDays(year, month))
}

func numberOfDays(year int, month int) int {
	m := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	if month != 2 {
		return m[month]
	} else {
		if year%4 == 0 {
			if year%100 == 0 {
				if year%400 != 0 {
					return m[month]
				}
			}
			return 29
		}
	}
	return m[month]
}

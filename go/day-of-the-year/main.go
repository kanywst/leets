package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// date := "2019-01-09"
	date := "2019-02-10"
	fmt.Println(dayOfYear(date))
}

func dayOfYear(date string) int {
	calender := map[string]int{
		"1":  31,
		"2":  30,
		"3":  31,
		"4":  30,
		"5":  31,
		"6":  30,
		"7":  31,
		"8":  31,
		"9":  30,
		"10": 31,
		"11": 30,
		"12": 31,
	}
	dd := strings.Split(date, "-")
	mm, _ := strconv.Atoi(dd[1])
	d, _ := strconv.Atoi(dd[2])
	cnt := 0
	fmt.Println(d)
	for i := 1; i < mm+1; i++ {
		m, _ := calender[string(i)]
		fmt.Println(m)
		cnt += m
	}
	cnt += d
	return cnt
}

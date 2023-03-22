package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	date1, date2 := "2020-01-15", "2019-12-31"
	fmt.Println(daysBetweenDates(date1, date2))
}

func daysBetweenDates(date1 string, date2 string) int {
	dt1, dt2 := strings.Split(date1, "-"), strings.Split(date2, "-")
	y1, _ := strconv.Atoi(dt1[0])
	m1, _ := strconv.Atoi(dt1[1])
	d1, _ := strconv.Atoi(dt1[2])
	y2, _ := strconv.Atoi(dt2[0])
	m2, _ := strconv.Atoi(dt2[1])
	d2, _ := strconv.Atoi(dt2[2])
	day1 := time.Date(y1, time.Month(m1), d1, 0, 0, 0, 0, time.Local)
	day2 := time.Date(y2, time.Month(m2), d2, 0, 0, 0, 0, time.Local)
	duration := day2.Sub(day1)
	hours0 := int(duration.Hours())
	days := hours0 / 24
	return int(math.Abs(float64(days)))
}

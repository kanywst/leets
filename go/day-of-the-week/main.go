package main

import (
	"fmt"
	"time"
)

func main() {
	day, month, year := 31, 8, 2019
	fmt.Println(dayOfTheWeek(day, month, year))
}

func dayOfTheWeek(day int, month int, year int) string {
	today := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	DayOfWeek := today.Weekday()
	return DayOfWeek.String()
}

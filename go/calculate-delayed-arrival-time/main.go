package main

import "fmt"

func main() {
	arrivalTime, delayedTime := 13, 11
	fmt.Println(findDelayedArrivalTime(arrivalTime, delayedTime))
}

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	return (arrivalTime + delayedTime) % 24
}

package main

import (
	"fmt"
	"math"
)

func main() {
	p1, p2, p3, p4 := []int{1, 0}, []int{0, 1}, []int{-1, 0}, []int{0, -1}
	fmt.Println(validSquare(p1, p2, p3, p4))
}

//p4p2
//p1p3

//p2
//  p1

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) (ans bool) {
	n := math.Abs(float64(p4[1] - p1[1]))
	fmt.Println(n, math.Abs(float64(p3[0]-p1[0])), math.Abs(float64(p2[1]-p3[1])), math.Abs(float64(p2[0]-p4[0])))
	if n == math.Abs(float64(p3[0]-p1[0])) && n == math.Abs(float64(p2[1]-p3[1])) && n == math.Abs(float64(p2[0]-p4[0])) {
		ans = true
	} else {
		ans = false
	}
	return
}

package main

import (
	"fmt"
)

func main() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) (ans bool) {
	x := coordinates[1][0] - coordinates[0][0]
	y := coordinates[1][1] - coordinates[0][1]
	inc := -1
	if x == 0 {
		inc = 0
	} else {
		inc = y / x
	}
	ans = true
	for i := 1; i < len(coordinates)-1; i++ {
		tmp_x := coordinates[i+1][0] - coordinates[i][0]
		tmp_y := coordinates[i+1][1] - coordinates[i][1]
		tmp_inc := -1
		if tmp_x == 0 {
			tmp_inc = 0
		} else {
			tmp_inc = tmp_y / tmp_x
		}
		if inc != tmp_inc {
			ans = false
			break
		}
	}
	return
}

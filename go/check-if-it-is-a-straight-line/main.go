package main

import "fmt"

func main() {
	// coordinates := [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}, []int{5, 6}, []int{6, 7}}
	// coordinates := [][]int{[]int{0, 0}, []int{0, 5}, []int{5, 5}, []int{5, 0}}
	coordinates := [][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}, []int{1, 5}, []int{1, 6}, []int{6, 7}}
	fmt.Println(checkStraightLine(coordinates))
}

func checkStraightLine(coordinates [][]int) bool {
	m := make(map[float64]int)
	for i := 0; i < len(coordinates)-1; i++ {
		y2_y1 := coordinates[i+1][1] - coordinates[i][1]
		x2_x1 := coordinates[i+1][0] - coordinates[i][0]
		if x2_x1 == 0 {
			m[0] += 1
			continue
		}
		if y2_y1 == 0 {
			m[1] += 1
			continue
		}
		m[float64(y2_y1)/float64(x2_x1)] += 1
	}
	if len(m) == 1 {
		return true
	}
	return false
}

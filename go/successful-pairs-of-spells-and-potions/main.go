package main

import "fmt"

func main() {
	spells, potions, success := []int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7
	fmt.Println(successfulPairs(spells, potions, int64(success)))
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	tmp := []int{}
	for _, i := range spells {
		cnt := 0
		for _, j := range potions {
			if j*i >= int(success) {
				cnt += 1
			}
		}
		tmp = append(tmp, cnt)
	}
	return tmp
}

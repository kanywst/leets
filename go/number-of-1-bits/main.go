package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}

func hammingWeight(num uint32) int {
	s := fmt.Sprintf("%b", num)
	cnt := 0
	for _, i := range string(s) {
		if string(i) == "1" {
			cnt += 1
		}
	}
	return cnt
}

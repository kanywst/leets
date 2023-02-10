package main

import "fmt"

func main() {
	n := "32"
	fmt.Println(minPartitions(n))
}
func minPartitions(n string) (ans int) {
	for _, c := range n {
		fmt.Println("c: ", c)
		fmt.Println("int(c-'0'): ", c-'0')
		if t := int(c - '0'); ans < t {
			ans = t
		}
	}
	return
}

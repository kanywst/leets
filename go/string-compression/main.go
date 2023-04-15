package main

import (
	"fmt"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Println(compress(chars))
}

func compress(chars []byte) (ans int) {
	m := make(map[byte]byte)
	for _, i := range chars {
		m[i] += 1
	}
	chars = []byte{}
	for k, v := range m {
		chars = append(chars, []byte{k, v}...)
		ans += len(string(k))
		ans += len(string(v))
	}
	fmt.Println(chars)
	return
}

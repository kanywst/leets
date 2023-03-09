package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10
	fmt.Println(countPrimes((n)))
}

func countPrimes(n int) (cnt int) {
	for _ = range Sieve(n) {
		cnt += 1
	}
	return
}

func Sieve(n int) chan int {
	s := make([]bool, n)
	for x := 2; x < int(math.Sqrt(float64(n)))+1; x++ {
		if !s[x] {
			for i := x + x; i < len(s); i = i + x {
				s[i] = true
			}
		}
	}

	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			if i > 1 && !s[i] {
				ch <- i
			}
		}
		close(ch)

	}()
	return ch

}

package main

import "fmt"

func permutation(n int, k int) int {
	v := 1
	fmt.Println("permutation:", n, k)
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			fmt.Println("n-i:", n-i)
			v *= (n - i)
			fmt.Println("v:", v)
		}
	} else if k > n {
		v = 0
	}
	fmt.Println("v:", v)
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	fmt.Println(permutation(n, k), factorial(k))
	return permutation(n, k) / factorial(k)
}

func main() {
	m, n := 13, 23
	fmt.Println(uniquePaths(m, n))
}

func uniquePaths(m int, n int) int {
	return combination(m-1+n-1, n-1)
}

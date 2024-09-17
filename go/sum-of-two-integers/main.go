package main

// func getSum(a int, b int) int {
//     return a+b
// }

func getSum(a int, b int) int {
	sum := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		a := sum
		b := carry
		sum = a ^ b
		carry = (a & b) << 1
	}
	return sum
}

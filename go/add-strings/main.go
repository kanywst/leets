package main

import "fmt"

func main() {
	num1, num2 := "11", "123"
	fmt.Println(addStrings(num1, num2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	maxLen := max(len1, len2)

	result := make([]byte, maxLen+1)
	carry := byte(0)

	for i := 0; i < maxLen || carry > 0; i++ {
		digit1 := byte(0)
		if i < len1 {
			digit1 = num1[len1-1-i] - '0'
		}

		digit2 := byte(0)
		if i < len2 {
			digit2 = num2[len2-1-i] - '0'
		}

		sum := digit1 + digit2 + carry
		result[maxLen-i] = sum%10 + '0'
		carry = sum / 10
	}

	if result[0] == 0 {
		result = result[1:]
	}

	return string(result)
}

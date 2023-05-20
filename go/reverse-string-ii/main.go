package main

import "fmt"

func main() {
	s, k := "hyzqyljrnigxvdtneasepfahmtyhlohwxmkqcdfehybknvdmfrfvtbsovjbdhevlfxpdaovjgunjqlimjkfnqcqnajmebeddqsgl", 39
	fmt.Println(reverseStr(s, k))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStr(s string, k int) (ans string) {
	i := 0
	if len(s) < 2 {
		ans = s
		return
	}
	if len(s) < k {
		ans = reverse(s)
		return
	}
	for {
		if i+2*k < len(s) {
			tmp := s[i : i+2*k]
			ans += reverse(tmp[:k]) + tmp[k:]
			i += 2 * k
		} else {
			tmp := s[i:]
			if len(tmp) < k {
				ans += reverse(tmp)
			} else {
				ans += reverse(tmp[:k]) + tmp[k:]
			}
			break
		}
	}
	return
}

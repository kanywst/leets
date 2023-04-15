package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	queryIP := "172.16.254.1"
	fmt.Println(validIPAddress(queryIP))
}

func validIPAddress(queryIP string) string {
	qq := strings.Split(".")
	if len(qq) == 4 {
		for _, i := range qq {
			ii, _ := strconv.Atoi(i)
			if ii >255
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00

	return []float64{kelvin, fahrenheit}
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	a := strings.Split(sc.Text(), " = ")
	celsius, _ := strconv.ParseFloat(a[1], 64)
	fmt.Println(convertTemperature(celsius))
}

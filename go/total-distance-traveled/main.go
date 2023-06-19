package main

import "fmt"

func main() {
	mainTank, additionalTank := 9, 2
	fmt.Println(distanceTraveled(mainTank, additionalTank))
}

func distanceTraveled(mainTank int, additionalTank int) (ans int) {
	for mainTank > 0 {
		if mainTank >= 5 {
			if additionalTank > 0 {
				additionalTank -= 1
				mainTank += 1
			}
			mainTank -= 5
			ans += 50
		} else {
			ans += mainTank * 10
			mainTank = 0
		}
	}
	return
}

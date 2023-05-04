package main

func monoInc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func monoDec(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}

func monoMountainInc(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func monoMountainDec(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= arr[i+1] {
			return false
		}
	}
	return true
}

func validMountainArray(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if monoInc(arr) || monoDec(arr) {
			return false
		}
		if len(arr[:i]) != 0 && len(arr[i:]) != 0 && monoMountainInc(arr[:i]) && monoMountainDec(arr[i:]) {
			return true
		}
	}
	return false
}

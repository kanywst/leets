package main

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	c, notZero := 0, []int{}
	for _, n := range nums {
		if n == 0 {
			c += 1
		} else {
			notZero = append(notZero, n)
		}
	}
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-c {
			nums[i] = notZero[i]
		} else {
			nums[i] = 0
		}
	}
	return
}

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
	}
	max := -32768
	for _, i := range nums {
		if max < i {
			max = i
		}
	}
	return max
}

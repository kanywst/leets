package main

func main() {

}

func isOdd(x []int, k int) bool {
	cnt := 0
	for _, i := range x {
		if i%2 == 1 {
			cnt += 1
		}
		if cnt == k {
			return true
		}
	}
	return false
}
func numberOfSubarrays(nums []int, k int) int {

}

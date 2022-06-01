package greedy

func canJump(nums []int) bool {
	max := nums[0]
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		if max < i {
			break
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= n {
			return true
		}
	}
	return false
}

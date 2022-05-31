package greedy

func jump(nums []int) int {
	var max, end, step int
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if i >= end {
			step++
			end = max
		}

	}
	return step
}

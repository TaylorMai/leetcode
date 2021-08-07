package greedy

/*
	需以三个连续的数考虑当前修改策略
	可优化
 */
func checkPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	var count int
	for i := 1; i < len(nums)-1; i++ {
		switch true {
		case nums[i-1] <= nums[i] && nums[i] <= nums[i+1]:
			continue
		case nums[i-1] <= nums[i] && nums[i] > nums[i+1]:
			if nums[i+1] >= nums[i-1] {
				nums[i] = nums[i-1]
			} else {
				nums[i+1] = nums[i]
			}
			count++
		case nums[i-1] > nums[i] && nums[i] < nums[i+1]:
			if nums[i-1] > nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i] = nums[i-1]
			}
			count++
		case nums[i-1] > nums[i] && nums[i] > nums[i+1]:
			return false
		case nums[i-1] > nums[i] && nums[i] == nums[i+1]:
			nums[i-1] = nums[i]
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
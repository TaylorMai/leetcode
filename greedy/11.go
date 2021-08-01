package greedy

/*
	首先面积=(x1-x2)*min(y1, y2)，其次需要能够枚举到全部情况。考虑到x1-x2的最大值为两边，可以通过移动两侧的高来尝试找到最大值。

	移动策略：因为每次向内移动，面积的底肯定会减少，为了找到最大值，需要高能够选到大的。因此每次应移动最小的高，这样才有可能选到后面的较大高。

	当两侧相同时：假设中间位置会有两个较大高，使得面积最大。此时其实无论你先移动哪一侧，由于每次总会淘汰较小高，因此两侧还是会遍历到两侧为上述假设的两条高的情况。
	（原始解的向里面探测其实可以不需要）
*/

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := 0
	for l < r {
		tmp := (r - l) * min(height[r], height[l])
		if tmp > res {
			res = tmp
		}
		switch true {
		case height[l] < height[r]:
			l++
		case height[l] > height[r]:
			r--
		default:
			if l+1 == r {
				l++ //random
				continue
			}
			if height[l+1] >= height[r-1] {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

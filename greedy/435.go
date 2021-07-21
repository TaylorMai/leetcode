package greedy

import (
	"math"
	"sort"
)

/*
	任意时候，优先选结尾更小的，将利于更多其他区间接在后面，从而令更多区间被用到、更少区间需要去除
 */

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1] || (intervals[i][1]==intervals[j][1] && intervals[i][0] < intervals[j][0])
	})
	var (
		del int
		end = math.MinInt32
	)
	for i := range intervals {
		if end <= intervals[i][0] {
			end = intervals[i][1]
		} else {
			del++
		}
	}
	return end
}

package greedy

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	str := make([]string, len(nums))
	sum := 0
	for i := range nums {
		v := strconv.FormatUint(uint64(nums[i]), 10)
		sum += len(v)
		str[i] = v
	}

	sort.Slice(str, func(i, j int) bool {
		if len(str[i]) == len(str[j]) {
			return str[i] > str[j]
		}
		v := str[i] + str[j]
		for p1, p2 := 0, len(str[i]); p1 < len(v); {
			if v[p1] != v[p2] {
				return v[p1] > v[p2]
			}
			p1++
			p2++
			p2 %= len(v)
		}
		return true
	})

	if str[0] == "0" {
		return "0"
	}

	var builder strings.Builder
	builder.Grow(sum)
	for i := range str {
		builder.WriteString(str[i])
	}
	return builder.String()

}

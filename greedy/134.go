package greedy

import "math"

func canCompleteCircuit(gas []int, cost []int) int {
	var sum, j int
	min := math.MaxInt
	for i := range gas {
		v := gas[i] - cost[i]
		sum += v
		if min >= sum {
			min = sum
			j = i
		}
	}
	if sum < 0 {
		return -1
	} else {
		return (j + 1) % len(gas)
	}
}

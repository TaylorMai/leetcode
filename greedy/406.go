package greedy

import "sort"

/*
	对于（h，k），意味着h前需要有k个>=h的数。假设该队列为降序，对于其中一个数，其后面的数字无论在哪里都不会影响他的k，而其前面的数字的位置则会产生影响。
 */

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || people[i][0] == people[j][0] && people[i][1] < people[j][1]
	})
	var ans [][]int
	for _, v := range people {
		k := v[1]                             //前面需要多少个比自己或等于自己高度的people
		tem := append([][]int{v}, ans[k:]...) //插入：先把后面的插进来
		ans = append(ans[:k], tem...)         //插入：再连带后面插入到正确的位置
	}
	return ans
}

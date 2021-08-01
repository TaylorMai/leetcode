package greedy

import (
	"math"
	"sort"
)

/*
	对于每个字符的范围，都希望找到一个可以刚好分开两个区间的点；因为需要分得越多越好，因此每次都先取符合条件的最相邻的点，因此需要先排序。
*/

func partitionLabels(s string) []int {
	//字符与数组下标映射
	dict := make(map[uint8]int)
	seq := 0
	for i := range s {
		if _, ok := dict[s[i]]; !ok {
			dict[s[i]] = seq
			seq++
		}
	}
	//转为[x,y]范围数组
	list := make([][2]int, len(dict))
	for i := range list {
		list[i][0] = math.MaxInt32
		list[i][1] = math.MinInt32
	}
	for i := range s {
		if i < list[dict[s[i]]][0] {
			list[dict[s[i]]][0] = i
		}
		if i > list[dict[s[i]]][1] {
			list[dict[s[i]]][1] = i
		}
	}
	//以起点排序
	sort.Slice(list, func(i, j int) bool {
		return list[i][0] < list[j][0]
	})
	//对于一个有序的范围，找到“上一个的结尾”与“下一个的起点”不重合的点
	res := make([]int, 0)
	front, end := 0, list[0][1]
	for i := range list {
		if end < list[i][0] { //不重合
			res = append(res, end-front+1)
			front = list[i][0]
		}
		if end < list[i][1] { //与后一个范围重叠了，需要融合两个范围，并以最大的结尾作为当前范围的结尾
			end = list[i][1]
		}
	}
	res = append(res, end-front+1)
	return res
}

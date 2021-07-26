package greedy

/*
	能种就种，前面和后面都没有就可以
 */

func canPlaceFlowers(flowerbed []int, n int) bool {
	var (
		done int
		pre  int
		l    = len(flowerbed) - 1
	)
	for i, v := range flowerbed {
		if v == 0 && pre == 0 && (i == l || flowerbed[i+1] == 0) {
			done++
			v = 1 //种花
		}
		pre = v
	}
	return done >= n
}

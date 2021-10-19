package other

func MyPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	res := float64(1)
	for {
		if n == 0 {
			break
		}
		if n&1 > 0 {
			res *= x
		}
		x *= x
		n = n >> 1
	}
	return res
}

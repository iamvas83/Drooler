package bitmanipulation

func Base2Decimal(A int, B int) int {

	ans := 0
	mul := 1
	for A > 0 {
		x := A % 10
		ans += mul * x
		mul *= B
		A /= 10
	}
	return ans
}

package bitmanipulation

func DecimalToAnyBase(A int, B int) int {

	ans := 0
	pow := 1
	for A > 0 {
		digit := A % B
		A /= B
		ans += digit * pow
		pow *= 10
	}
	return ans
}

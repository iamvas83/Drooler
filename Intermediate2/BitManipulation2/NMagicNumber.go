package bitmanipulation2

func MagicNumberNth(A int) int {
	ans := 0
	pow := 5
	for A > 0 {
		r := A % 2
		A = A / 2
		ans += (r * pow)
		pow *= 5
	}
	return ans
}

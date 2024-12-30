package bitmanipulation2

/*
Each bit in the binary representation of A corresponds to a power of 5.
The rightmost bit corresponds to 5^1,
the next to 5^2, then 5^3, and so on.
If a bit is 1, we include that power of 5 in our sum.
*/
func MagicNumberNth(A int) int {
	ans := 0
	pow := 5
	for A > 0 {
		r := A % 2
		//A = A / 2
		A = A >> 1
		ans += (r * pow)
		pow *= 5
	}
	return ans
}

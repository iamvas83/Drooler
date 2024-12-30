package bitmanipulation2

func ReverseBits2(A uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans <<= 1
		if (A & (1 << i)) != 0 {
			ans |= 1
		}

	}
	return ans
}

/* It iterates 32 times (for each bit)
Left-shifts the answer to make room for the new bit
Checks if the rightmost bit of A is 1
If it is, it sets the rightmost bit of ans to 1
Right-shifts A to move to the next bit
This effectively reverses the bits of the input number.
*/

func ReverseBits(A uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans <<= 1

		if A&1 == 1 {
			ans = ans | 1
		}
		A = A >> 1
	}
	return ans
}

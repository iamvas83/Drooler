package bitmanipulation2

func ReverseBits(A uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		ans <<= 1
		if (A & (1 << i)) != 0 {
			ans |= 1
		}

	}
	return ans
}

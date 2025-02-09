package recirsion2

func KthSymbol(A int, K int64) int {
	B := int(K)

	if B == 0 {
		return 0
	}

	val := KthSymbol(A-1, int64(B/2))

	if B%2 == 0 {
		return val
	} else {
		return 1 - val
	}
}

package gcd

func Gcd(A, B int) int {
	if A == 0 {
		return B
	}

	return Gcd(B%A, A)
}

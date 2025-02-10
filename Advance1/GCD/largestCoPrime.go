package gcd

func LargestCoPrimeFactor(A, B int) int {
	x := 0
	for A > 0 {
		res := gGCDcd(A, B)
		A = A / res
		x = A
		if res == 1 {
			break
		}
	}
	return x
}

func gGCDcd(A, B int) int {
	if A == 0 {
		return B
	}

	return gcd(B%A, A)
}

package modulararithmatic

func LargePow(A int, B int) int {
	var fact int = 1
	var mod int = 1e9 + 7
	for i := 2; i <= B; i++ {
		fact = (fact * i) % (mod - 1)
	}
	ans := pow(A, fact, mod)
	return ans
}

func pow(A, B, mod int) int {
	var ans int = 1

	for B > 0 {
		if B%2 != 0 {
			ans = (ans * A) % mod
		}
		A = ((A % mod) * (A % mod)) % mod
		B /= 2
	}
	return ans % mod
}

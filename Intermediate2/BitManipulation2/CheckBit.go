package bitmanipulation2

func CheckBit(A, B int) int {
	mask := 1 << B
	ans := A & mask
	if ans == 1 {
		return 1
	} else {
		return 0
	}
}

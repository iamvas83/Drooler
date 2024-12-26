package bitmanipulation2

//Set Bit -> 0->1 or 1-> 1
func SetBit0(A, B int) int {
	res := 0
	res = res | 1<<A
	res = res | 1<<B
	return res
}

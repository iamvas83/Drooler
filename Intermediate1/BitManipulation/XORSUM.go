package bitmanipulation

func XOR(A int, B int) int {
	x := A & B
	/*
			1011
		&   0101
		________
			0001

			1011
		^	0001
		--------
			1010

			0101
		^	0001
		--------
			0100

			1010
		+	0100
		--------
			1110
	*/

	return A ^ x + B ^ x
}

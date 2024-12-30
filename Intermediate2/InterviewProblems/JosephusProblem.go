package interviewproblems

import "math"

func JosephusProblem(A int) int {
	binx := math.Log2(float64(A))
	xint := int(binx)
	pow := math.Pow(2, float64(xint))
	return 2*(A-int(pow)) + 1
}

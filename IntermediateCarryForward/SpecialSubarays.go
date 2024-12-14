package intermediatecarryforward

/**
 * @input A : String
 *
 * @Output Integer
 */

func SpecialSubArrays(as string) int {

	n := len(as)
	cnt_c := 0
	count := 0
	for i := n - 1; i >= 0; i-- {
		if as[i] == 'A' || as[i] == 'E' || as[i] == 'I' || as[i] == 'O' || as[i] == 'U' || as[i] == 'a' || as[i] == 'e' || as[i] == 'i' || as[i] == 'o' || as[i] == 'u' {
			cnt_c++
			count += cnt_c
		} else {
			cnt_c++
		}
	}
	return count % 10003
}

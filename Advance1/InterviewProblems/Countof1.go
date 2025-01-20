package arrays3

func CountOf1(A int) int {
	count := 0
	multiplier := 1

	for A/multiplier > 0 {
		// Extract higher, current, and lower digits
		lower := A % multiplier
		current := (A / multiplier) % 10
		higher := A / (multiplier * 10)

		// Count the number of ones contributed by the current digit
		if current == 0 {
			count += higher * multiplier
		} else if current == 1 {
			count += higher*multiplier + lower + 1
		} else {
			count += (higher + 1) * multiplier
		}

		// Move to the next digit position
		multiplier *= 10
	}

	return count
}

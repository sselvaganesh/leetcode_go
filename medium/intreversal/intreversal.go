package intreversal

func reverse(x int) int {

	var absVal, target, rem int

	if x < 0 {
		absVal = x * -1
	} else {
		absVal = x
	}

	for i := 1; absVal > 0; i++ {
		rem = absVal % 10
		absVal = absVal / 10

		target = (target * 10) + rem
		if target > math.MaxInt32 || target < math.MinInt32 {
			return 0
		}

	}

	if x < 0 {
		return -1 * target
	}

	return target
}

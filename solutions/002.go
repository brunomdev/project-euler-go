package solutions

func fibonacci() func() int {
	x, y := 0, 1

	return func() int {
		x, y = y, x+y
		return y
	}
}

// Solution002 return the soltion of problem 002 from Project Euler
func Solution002(limit int) int {
	sum := 0

	f := fibonacci()
	for sum < limit {
		x := f()

		if x%2 == 0 {
			sum += x
		}
	}

	return sum
}

package solutions

// Solution001 return the soltion of problem 001 from Project Euler
func Solution001(limit int) int {
	sum := 0

	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

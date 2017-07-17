package solutions

// Solution003 return the solution of problem 003 from Project Euler
func Solution003(n int) int {
	// Verifies if number is divisible by 2
	for n%2 == 0 {
		n = n / 2
	}

	// Performs the division of n by odd numbers
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			n = n / i
		}
	}

	return n
}

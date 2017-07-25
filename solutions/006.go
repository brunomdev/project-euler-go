package solutions

// Solution006 return the solution of problem 006 from Project Euler
func Solution006(n int) int {
	return sumNaturalSquares(n) - sumSquares(n)
}

func sumSquares(limit int) (sum int) {
	for i := 1; i <= limit; i++ {
		sum += i * i
	}

	return
}

func sumNaturalSquares(limit int) (sum int) {
	for i := 1; i <= limit; i++ {
		sum += i
	}

	sum = sum * sum

	return
}

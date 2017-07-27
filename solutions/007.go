package solutions

// Solution007 return the solution of problem 007 from Project Euler
func Solution007(n int) int {
	p := 1
	var primes []int

	for i := 1; p <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)

			p++
		}
	}

	return primes[n-1]
}

func isPrime(value int) bool {
	if value <= 1 {
		return false
	}

	for i := 2; i <= value/2; i++ {
		if value%i == 0 {
			return false
		}
	}

	return true
}

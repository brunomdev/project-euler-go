package solutions

// Solution005 return the solution of problem 005 from Project Euler
func Solution005(n int64) int64 {
	var a, i int64 = 1, 2

	for ; i <= n; i++ {
		a = a * i / gcd(a, i)
	}

	return a
}

func gcd(a, b int64) int64 {
	c := a % b

	if c == 0 {
		return b
	}

	return gcd(b, c)
}

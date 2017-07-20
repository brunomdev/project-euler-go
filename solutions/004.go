package solutions

import "strconv"

// Solution004 return the solution of problem 004 from Project Euler
func Solution004(n int) int {
	palindrome := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			mNumber := j * i

			if isPalindrome(mNumber) && mNumber > palindrome {
				palindrome = mNumber
			}
		}
	}

	return palindrome
}

func isPalindrome(n int) bool {
	if strconv.Itoa(n) == Reverse(strconv.Itoa(n)) {
		return true
	}

	return false
}

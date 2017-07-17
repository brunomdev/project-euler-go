package solutions

import "testing"

var testsSolution003 = []testData{
	{13195, 29},
	{600851475143, 6857},
}

func TestSolution003(t *testing.T) {
	for _, test := range testsSolution003 {
		if sum := Solution003(test.limit); sum != test.expected {
			t.Errorf("Solution003(%d) = %d, should be %d\n", test.limit, sum, test.expected)
		}
	}
}

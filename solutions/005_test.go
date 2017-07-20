package solutions

import "testing"

var testsSolution005 = []testData{
	{10, 2520},
	{20, 232792560},
}

func TestSolution005(t *testing.T) {
	for _, test := range testsSolution005 {
		if sum := Solution005(int64(test.in)); sum != int64(test.expected) {
			t.Errorf("Solution005(%d) = %d, should be %d\n", test.in, sum, test.expected)
		}
	}
}

package solutions

import "testing"

var testsSolution007 = []testData{
	{6, 13},
	{10001, 104743},
}

func TestSolution007(t *testing.T) {
	for _, test := range testsSolution007 {
		if sum := Solution007(test.in); sum != test.expected {
			t.Errorf("Solution007(%d) = %d, should be %d\n", test.in, sum, test.expected)
		}
	}
}

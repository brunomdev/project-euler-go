package solutions

import "testing"

var testsSolution004 = []testData{
	{99, 9009},
	{999, 906609},
}

func TestSolution004(t *testing.T) {
	for _, test := range testsSolution004 {
		if sum := Solution004(test.in); sum != test.expected {
			t.Errorf("Solution004(%d) = %d, should be %d\n", test.in, sum, test.expected)
		}
	}
}

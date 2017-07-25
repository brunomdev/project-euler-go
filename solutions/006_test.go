package solutions

import "testing"

var testsSolution006 = []testData{
	{10, 2640},
	{100, 25164150},
}

func TestSolution006(t *testing.T) {
	for _, test := range testsSolution006 {
		if sum := Solution006(test.in); sum != test.expected {
			t.Errorf("Solution006(%d) = %d, should be %d\n", test.in, sum, test.expected)
		}
	}
}

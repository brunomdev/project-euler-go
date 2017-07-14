package solutions

import "testing"

var testsSolution001 = []testData{
	{10, 23},
	{1000, 233168},
}

func TestSolution001(t *testing.T) {
	for _, test := range testsSolution001 {
		if sum := Solution001(test.limit); sum != test.expected {
			t.Errorf("Solution001(%d) = %d, should be %d\n", test.limit, sum, test.expected)
		}
	}

}

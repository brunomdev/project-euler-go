package solutions

import "testing"

var testsSolution002 = []testData{
	{4000000, 4613732},
}

func TestSolution002(t *testing.T) {
	for _, test := range testsSolution002 {
		if sum := Solution002(test.limit); sum != test.expected {
			t.Errorf("Solution002(%d) = %d, should be %d\n", test.limit, sum, test.expected)
		}
	}
}

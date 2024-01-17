package goquest

import "testing"

func TestAdd(t *testing.T) {
	type addTest struct {
		a, b, expected int
	}

	tests := []addTest{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}

	for _, test := range tests {
		if output := Add(test.a, test.b); output != test.expected {
			t.Errorf("Add %d + %d expected %d. Got = %d", test.a, test.b, output, test.expected)
		}
	}
}

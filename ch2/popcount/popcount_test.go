package popcount

import "testing"

func TestPopCount(t *testing.T) {
	var tests = []struct {
		num uint64
		expected int
	}{
		{28867, 7},
		{55573, 8},
		{10012002, 10},
	}
	for _, test := range tests {
		if actual := PopCount(test.num); actual != test.expected {
			t.Errorf("expected %d but got %d", test.expected, actual)
		}
	}
}

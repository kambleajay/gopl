package popcount

import "testing"

var tests = []struct {
	num uint64
	expected int
}{
	{28867, 7},
	{55573, 8},
	{10012002, 10},
}

func TestPopCount(t *testing.T) {
	for _, test := range tests {
		if actual := PopCount(test.num); actual != test.expected {
			t.Errorf("expected %d but got %d", test.expected, actual)
		}
	}
}

func TestPopCountLoop(t *testing.T) {
	for _, test := range tests {
		if actual := PopCountLoop(test.num); actual != test.expected {
			t.Errorf("expected %d but got %d", test.expected, actual)
		}
	}
}

//result: BenchmarkPopCount-8		2000000000			 0.30 ns/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(10012002)
	}
}

//result: BenchmarkPopCountLoop-8		100000000			20.1 ns/op
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(10012002)
	}
}

package emptyinterface

import "testing"

var tests = []struct {
	input  int
	output interface{}
}{
	{0, interface{}(42)},
	{1, interface{}("Jesus!")},
	{2, interface{}(struct {
		x int
		y int
	}{3, -1})},
	{3, "FAIL"},
}

func TestEmptyInterface(t *testing.T) {
	for _, test := range tests {
		out := EmptyInterface(test.input)
		if out != test.output {
			t.Errorf("Input %d yielded the incorrect result.", test.input)
		}
	}
}

func BenchmarkEmptyInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			EmptyInterface(test.input)
		}
	}
}

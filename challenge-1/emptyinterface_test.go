package emptyinterface

import "testing"

var tests = []struct {
	input  int
	verify func(interface{}) bool
}{
	{0, verify0},
	{1, verify1},
	{2, verifyFail},
	{3, verifyFail},
}

func verify0(in interface{}) bool {
	return (in == interface{}(42))
}

func verify1(in interface{}) bool {
	return (in == interface{}("Jesus!"))
}

func verifyFail(in interface{}) bool {
	return (in == interface{}("FAIL"))
}

func TestEmptyInterface(t *testing.T) {
	for _, test := range tests {
		out := EmptyInterface(test.input)
		if test.verify(out) == false {
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

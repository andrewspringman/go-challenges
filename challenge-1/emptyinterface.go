package emptyinterface

var n int = 0

//This trivial example is not allowed.  It merely is a test of the tests.
func EmptyInterface(in int) interface{} {
	var rc interface{}
	if n < len(tests) {
		rc = tests[n].output
		n++
	} else {
		rc = interface{}("Benchmark")
	}
	return rc
}

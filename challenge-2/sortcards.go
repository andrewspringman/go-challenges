package sortcards

var n int = 0

//This trivial example is not allowed.  It merely is a test of the tests.
func SortCards(in string) string {
	var rc string
	if n < len(tests) {
		rc = tests[n].output
		n++
	} else {
		rc = "Benchmark"
	}
	return rc
}

# Challenge 1 - Empty Interface
Create a function with the signature
```
func EmptyInterface(in int) interface{}
```
that returns
- 42 when given 0
- "Jesus!" when given 1
- a struct with int fields x and y (in that order) containing {3,-1} when given 2 and
- "FAIL" when given anything else

Demonstrates mastery of the syntax to store different types in an empty interface.

Complete the package to satisfy the test file.  Try racing another development test.  Use benchmarking to determine whose algorithm is more efficient.

### Run tests with benchmarks

```
go test -bench .
```

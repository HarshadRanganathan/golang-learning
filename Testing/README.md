[Testing](#testing)

- [go test](#go-test)

  - [Examples](#examples)

  - [Benchmark](#benchmark)

  - [Coverage](#coverage)



## Testing

Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the “go test” command, which automates execution of any function of the form

``func TestXxx(*testing.T)``

where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.

To write a new test suite, create a file whose name ends ``_test.go`` that contains the ``TestXxx functions`` as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the “go test” command is run.

Let's consider below function for which we will write our tests.

```go
// Package hello provides function for division operation
package hello

// Divide takes in two operands and returns a result
func Divide(x int, y int) int {
  if y == 0 {
    panic("Divide by zero")
  }
  return x / y
}
```

Test file is named as `hello_test.go`

```go
package hello

import "testing"

func TestDivide(t *testing.T) {
  x := 10
  y := 2
  res := Divide(x, y)
  if res != 5 {
    t.Errorf("Div(%d/%d) = %d; want 5", x, y, res)
  }
}
```

### go test

go test command will look for any tests in any of the files in the current folder and run them.

```shell
$ go test
PASS
ok      _/C_/GoWorkspace/test/src/hello 1.126s
```

Incase of test failure,

```shell
$ go test
--- FAIL: TestDivide (0.00s)
    hello_test.go:10: Div(10/2) = 3; want 5
FAIL
exit status 1
FAIL    _/C_/GoWorkspace/test/src/hello 1.326s
```

#### Examples

The package also runs and verifies example code. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run.

```go
func ExampleDivide() {
  fmt.Println(Divide(10, 2))
  // Output: 5
}
```

Incase of example failure,

```shell
$ go test
--- FAIL: ExampleDivide (0.00s)
got:

want:
5
FAIL
exit status 1
FAIL    _/C_/GoWorkspace/test/src/hello 1.130s
```

#### Benchmark

Functions of the form

``func BenchmarkXxx(*testing.B)``

are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

```go
func BenchmarkDivide(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Divide(10, 2)
  }
}
```

Output:

```shell
$ go test -bench .
goos: windows
goarch: amd64
pkg: hello
BenchmarkDivide-12      2000000000               0.25 ns/op
PASS
ok      hello   3.032s
```

#### Coverage

'go test -cover' rewrites the source code with annotations to track which parts of each function are executed. 

It operates on one Go source file at a time, computing approximate basic block information by studying the source.

```shell
$ go test -cover .
ok      hello   0.511s  coverage: 66.7% of statements
```
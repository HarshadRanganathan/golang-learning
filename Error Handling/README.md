[Error Handling](#error-handling)

- [Return Errors Using Built-in Functions](#return-errors-using-built-in-functions)

- [Custom Errors](#custom-errors)

- [Panic](#Panic)

- [Recover](#Recover)

## Error Handling

In Go, functions may return values of type `error`. We need to check if the function has returned an error and take appropriate action.

```go
package main

import (
  "log"
  "os"
)

func main() {
  // Open, if successful, methods on returned file can be used for reading, otherwise, returns an error
  f, err := os.Open("")
  if err != nil {
    log.Fatal(err) // print err followed by call to os.Exit(1)
  }
  defer f.Close()
}
```

Output:

```text
2019/11/03 18:00:30 open : The system cannot find the file specified.
exit status 1
```

### Return Errors Using Built-in Functions

`errors.New` built-in function can be used to build a custom error message.

Go allows you to write functions that can return more than one result. Below example shows a function that returns a value if it completes successfully along with a potential error if incase it fails.

```go
package main

import (
  "errors"
  "fmt"
  "log"
)

func divide(x int, y int) (int, error) {
  if y == 0 {
    return -1, errors.New("Division by zero")
  }
  return x / y, nil
}

func main() {
  res, err := divide(10, 0)
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println(res)
  }
}
```

Output:

```text
2019/11/03 18:29:33 Division by zero
exit status 1
```

`fmt.Errorf` is another built-in function that helps you to build dynamic error messages using a format specifier.

```go
package main

import (
  "fmt"
  "log"
)

func divide(x int, y int) (int, error) {
  if y == 0 {
    return -1, fmt.Errorf("Division by zero: %v/%v", x, y)
  }
  return x / y, nil
}

func main() {
  res, err := divide(10, 0)
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println(res)
  }
}
```

Output:

```text
2019/11/03 18:39:11 Division by zero: 10/0
exit status 1
```

### Custom Errors

You can make use of `errors.New` and `fmt.Errorf` built-in functions to create custom errors.

If you need to create errors with more information, then implement the `error` interface.

```go
type error interface {
    Error() string
}
```

We'll follow the approach used in standard library `errors.New` which implements the error interface with a struct type to provide more information about the error.

```go
package main

import (
  "fmt"
  "log"
)

type divideByZero struct {
  x   int
  y   int
  err error
}

func (d divideByZero) Error() string {
  return fmt.Sprintf("Divide by zero error occurred: %v %v %v", d.x, d.y, d.err)
}

func divide(x int, y int) (int, error) {
  if y == 0 {
    err := fmt.Errorf("Division by zero: %v/%v", x, y)
    return -1, divideByZero{x, y, err}
  }
  return x / y, nil
}

func main() {
  res, err := divide(10, 0)
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println(res)
  }
}
```

### Panic

Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. 

Panics can be caused by runtime errors.

```go
package main

import (
  "fmt"
)

func divide(x int, y int) int {
  return x / y
}

func main() {
  res := divide(10, 0)
  fmt.Println(res)
}
```

Output:

```text
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.divide(...)
        C:/GoWorkspace/test/hello.go:8
main.main()
        C:/GoWorkspace/test/hello.go:12 +0x12
exit status 2
```

Panic can also be initiated by invoking ``panic`` directly.

```go
package main

import (
  "fmt"
)

func divide(x int, y int) int {
  if y == 0 {
    panic("Divide by zero")
  }
  return x / y
}

func main() {
  res := divide(10, 0)
  fmt.Println(res)
}
```

Output:

```text
panic: Divide by zero

goroutine 1 [running]:
main.divide(...)
        C:/GoWorkspace/test/hello.go:9
main.main()
        C:/GoWorkspace/test/hello.go:15 +0x41
exit status 2
```

### Recover

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. 

If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

```go
package main

import (
  "fmt"
)

func divide(x int, y int) int {
  if y == 0 {
    panic("Divide by zero")
  }
  return x / y
}

func main() {
  defer func() {
    if r := recover(); r != nil {
      fmt.Println("Divide by zero error occurred, recovering function execution, error:", r)
    }
  }()
  res := divide(10, 0)
  fmt.Println(res)
}
```

Output:

```text
Divide by zero error occurred, recovering function execution, error: Divide by zero
```
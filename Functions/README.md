[Functions](#functions)

- [Variadic Parameter](#variadic-parameter)

- [Defer Statements](#defer-statements)

- [Method Declarations](#method-declarations)

## Functions

A function can take zero or more arguments.

```go
package main

import "fmt"

func functionWithNoParameters() {
  fmt.Println("Function with no parameters called")
}

func functionWithParameters(s string) {
  fmt.Println(s)
}

func functionWithParametersAndResult(s string) string {
  return s
}

func main() {
  functionWithNoParameters()

  functionWithParameters("Function with parameters called")

  res := functionWithParametersAndResult("Function with parameters and result called")
  fmt.Println(res)
}
```

### Variadic Parameter

The final incoming parameter in a function signature may have a type prefixed with `...`. A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.

The value passed is a new slice of type []T with a new underlying array.

```go
package main

import "fmt"

func functionWithVariadicParameters(x ...int) {
  fmt.Println("Function with variadic parameters called")
  fmt.Println(x) // [1 2 3]
}

func main() {
  functionWithVariadicParameters(1, 2, 3)
}
```

If the final argument is a slice, it may be passed unchanged as the value for a ...T parameter if the argument is followed by `...`. In this case no new slice is created.

```go
package main

import "fmt"

func functionWithVariadicParameters(x ...int) {
  fmt.Println("Function with variadic parameters called")
  fmt.Println(x) // [1 2 3]
}

func main() {
  x := []int{1, 2, 3}
  functionWithVariadicParameters(x...)
}
```

### Defer Statements

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.

```go
package main

import "fmt"

func f1() {
  fmt.Println("f1")
}

func f2() {
  fmt.Println("f2")
}

func main() {
  defer f1()
  f2()
}
```

Output:

```text
f2
f1
```

### Method Declarations

A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

The receiver is specified via an extra parameter section preceding the method name.

```go
package main

import "fmt"

type person struct {
  firstName string
  lastName  string
}

/*
  Function with a receiver of type person
*/
func (p person) name() string {
  return p.firstName + " " + p.lastName
}

func main() {
  p := person{
    firstName: "James",
    lastName:  "Bond",
  }

  // method 'name' is bound to the receiver of type 'person' and is visible only within selectors for type 'person'
  fmt.Println(p.name())
}
```

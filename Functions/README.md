[Functions](#functions)

- [Variadic Parameter](#variadic-parameter)

- [Defer Statements](#defer-statements)

- [Method Declarations](#method-declarations)

- [Anonymous Functions](#anonymous-functions)

- [Closures](#closures)

- [Recursion](#recursion)

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

A function can return another function.

```go
package main

import "fmt"

func main() {
  f := echo()
  fmt.Printf("%T\n", f)
  fmt.Println(f("James bond")) // function invocation
}

// echo returns an anonymous function
func echo() func(string) string {
  return func(x string) string {
    return x
  }
}
```

Output:

```text
func(string) string
James bond
```

You can also pass function as arguments to another function and invoke it.

```go
package main

import "fmt"

func sum(x ...int) int {
  total := 0
  for _, v := range x {
    total += v
  }
  return total
}

func evenSum(f func(num ...int) int, x ...int) int {
  var y []int
  for _, v := range x {
    if v%2 == 0 {
      y = append(y, v)
    }
  }
  return f(y...)
}

func main() {
  num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

  sumOfAllNumbers := sum(num...)
  fmt.Printf("Sum of all numbers: %v\n", sumOfAllNumbers)

  sumOfEvenNumbers := evenSum(sum, num...)
  fmt.Printf("Sum of even numbers: %v", sumOfEvenNumbers)
}
```

Output:

```text
Sum of all numbers: 45
Sum of even numbers: 20
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

### Anonymous Functions

A function literal represents an anonymous function.

```go
package main

import "fmt"

func main() {
  // Anonymous function with arguments
  func(x string) {
    fmt.Println(x)
  }("Anonymous Function")
}
```

A function literal can be assigned to a variable or invoked directly.

```go
package main

import "fmt"

func main() {
  f := func(x string) {
    fmt.Println(x)
  }
  f("Anonymous Function")
}
```

### Closures

Function literals are closures: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.

```go
package main

import "fmt"

func incrementor() func() int {
  var x int
  return func() int {
    x++
    return x
  }
}

func main() {
  a := incrementor() // closure returned with x as 0
  b := incrementor() // closure returned with x as 0
  fmt.Println(a())   // increments x shared with closure assigned to 'a'
  fmt.Println(a())
  fmt.Println(a())
  fmt.Println(b()) // increments x shared with closure assigned to 'b'
  fmt.Println(b())
  fmt.Println(b())
}
```

Output:

```text
1
2
3
1
2
3
```

### Recursion

Call a function from within its own code.

```go
package main

import "fmt"

func factorial(n int) int {
  if n == 1 {
    return 1
  }
  return n * factorial(n-1)
}

func main() {
  f := factorial(5)
  fmt.Println(f) // 120
}
```

[Variable declarations](#variable-declarations)

[Zero value](#zero-value)

[Short variable declarations](#short-variable-declarations)

[Custom types](#custom-types)

[Type conversion](#type-conversion)

[String formatting](#string-formatting)

## Variable declarations

A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value.

You should use `var` keyword to declare variables.

```go
package main

import "fmt"

// variables are declared at package scope
var x int
var y string
var z bool

func main() {
  fmt.Println(x, y, z)
}
```

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment.

## Zero value

When storage is allocated for a variable, either through a declaration or a call of new, or when a new value is created, either through a composite literal or a call of make, and no explicit initialization is provided, the variable or value is given a default value.

```go
var x int // default 0
var y string // default ""
var z bool // default false
```

## Short variable declarations

Shorthand for a regular variable declaration with initializer expressions but no types.

```go
package main

import "fmt"

func main() {
  x := 42
  y := "James Bond"
  z := true

  fmt.Println(x, y, z)
}
```

Here, `x := 42` is a shorthand for variable declaration `var x = 42`.

In some contexts such as the initializers for "if", "for", or "switch" statements, they can be used to declare local temporary variables.

## Custom types

```go
package main

import "fmt"

type golang int // custom type golang with underlying type as int

var x golang

func main() {
  fmt.Println(x) // 0
  fmt.Printf("%T\n", x) // main.golang
  x = 42
  fmt.Println(x) // 42
}
```

## Type conversion

```go
package main

import "fmt"

type golang int

var x golang
var y int

func main() {
  fmt.Println(x) // 0
  fmt.Printf("%T\n", x) // main.golang
  x = 42
  fmt.Println(x) // 42
  y = int(x)
  fmt.Printf("%T\n", y) // int
  fmt.Println(y) // 42
}
```

## String formatting

Package `fmt` implements formatted I/O with functions analogous to C's printf and scanf.

`Sprintf` default behavior is for each formatting verb to format successive arguments passed in the call.

```go
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
  s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
  fmt.Println(s) // 42      James Bond      true
}
```

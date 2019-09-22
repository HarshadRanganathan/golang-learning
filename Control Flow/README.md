[For statements](#for-statements)

- [While condition](#while-condition)

- [Range clause](#range-clause)

[Break statements](#break-statements)

[Continue statements](#continue-statements)

[If statements](#if-statements)

- [Simple statement with condition](#simple-statement-with-condition)

[Switch statements](#switch-statements)

- [Type assertion](#type-assertion)

[Fallthrough statements](#fallthrough-statements)

## For statements

A "for" statement with a "for" clause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement.

```go
package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    fmt.Println(i)
  }
}
```

### While condition

For is Go's "while".

```go
package main

import "fmt"

func main() {
  i := 0
  for i <= 10 {
    fmt.Println(i)
    i++
  }
}
```

### Range clause

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel.

```go
package main

import "fmt"

func main() {
  for i, v := range "hello" {
    fmt.Printf("%v %v\n", i, v) // prints index and rune (ascii)
  }
}
```

Output:

```text
0 104
1 101
2 108
3 108
4 111
```

## Break statements

A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function.

```go
package main

import "fmt"

func main() {
  i := 0
  for {
    if i > 10 { break }
    fmt.Println(i)
    i++
  }
}
```

## Continue statements

A "continue" statement begins the next iteration of the innermost "for" loop at its post statement.

```go
package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    if i%2 == 0 {
      continue
    }
    fmt.Println(i) // Prints 1 3 5 7 9
  }
}
```

## If statements

If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.

```go
package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    if i%2 == 0 {
      fmt.Printf("%v is an even number\n", i)
    } else {
      fmt.Printf("%v is an odd number\n", i)
    }
  }
}
```

### Simple statement with condition

The expression may be preceded by a simple statement, which executes before the expression is evaluated.

```go
package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    if n := 2; i%n == 0 {
      fmt.Printf("%v is an even number\n", i)
    } else {
      fmt.Printf("%v is an odd number\n", i)
    }
  }
}
```

## Switch statements

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case.

If no case matches and there is a "default" case, its statements are executed.

```go
package main

import "fmt"

func main() {
  i := 2
  switch {
  case i == 2:
    fmt.Println("Value is equal to 2")
  default:
    fmt.Println("Value is not equal to 2")
  }

  switch "bananas" {
  case "Oranges", "Limes":
    fmt.Println("Citrus fruit")
  case "apricots", "peaches", "plums":
    fmt.Println("Stone fruit")
  case "mangoes", "bananas":
    fmt.Println("Topical fruit")
  }
}

```

### Type assertion

Type assertions can be done in switch statements.

```go
package main

import "fmt"

func main() {
  var i interface{} = 2
  switch i.(type) {
  case int:
    fmt.Println("Value is of type integer")
  default:
    fmt.Println("Unknown type")
  }
}
```

Output:

```text
Value is of type integer
```

## Fallthrough statements

Statement to indicate that control should flow from the end of this clause to the first statement of the next clause.

```go
package main

import "fmt"

func main() {
  switch "bananas" {
  case "Oranges", "Limes":
    fmt.Println("Citrus fruit")
    fallthrough
  case "apricots", "peaches", "plums":
    fmt.Println("Stone fruit")
    fallthrough
  case "mangoes", "bananas":
    fmt.Println("Topical fruit")
    fallthrough
  default:
    fmt.Println("Fruits are a good source of vitamins and minerals")
  }
}
```

Output:

```text
Topical fruit
Fruits are a good source of vitamins and minerals
```

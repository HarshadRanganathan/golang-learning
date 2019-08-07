[For statements](#for-statements)

[Break statements](#break-statements)

[Continue statements](#continue-statements)

[If statements](#if-statements)

[Switch statements](#switch-statements)

[Fallthrough statements](#fallthrough-statements)

## For statements

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an init and a post statement, such as an assignment, an increment or decrement statement.

```go
package main

import "fmt"

func main() {
  for i := 0; i <= 10; i++ {
    fmt.Println(i)
  }
}
```

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

```
Topical fruit
Fruits are a good source of vitamins and minerals
```

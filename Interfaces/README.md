[Interfaces](#Interfaces)

## Interfaces

An interface type specifies a method set called its interface.

More than one type may implement an interface.

```go
package main

import "fmt"

type person struct {
  firstName string
  lastName  string
}

// interface with type declaration
type human interface {
  speak()
}

func (p person) name() string {
  return p.firstName + " " + p.lastName
}

// P implements human interface
func (p person) speak() string {
  return "I am " + p.firstName + " " + p.lastName
}

func main() {
  p := person{
    firstName: "James",
    lastName:  "Bond",
  }

  fmt.Println(p.name())
  fmt.Println(p.speak())
}
```

Output:

```text
James Bond
I am James Bond
```

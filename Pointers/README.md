[Pointers](#pointers)

- [Selectors](#selectors)

## Pointers

For an operand x of type T, the address operation &x generates a pointer of type *T to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation.

For an operand x of pointer type *T, the pointer indirection *x denotes the variable of type T pointed to by x. If x is nil, an attempt to evaluate *x will cause a run-time panic.

```go
package main

import "fmt"

func main() {
  a := 10
  fmt.Println(a)         // value of 'a'
  fmt.Println(&a)        // address of 'a'
  fmt.Printf("%T\n", &a) // address operation &a generates a pointer of type *int to 'a'

  b := &a // assign address of 'a' to 'b'
  fmt.Println(b)
  fmt.Println(*b) // pointer indirection *b to variable 'a'

  *b = 20 // update value pointed to by pointer indirection *b to variable 'a'
  fmt.Println(a)
}
```

### Selectors

If the type of x is a defined pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.

```go
package main

import "fmt"

type person struct {
  firstName string
  lastName  string
}

func name(p *person) {
  // p.firstName is a shorthand for (*p).firstName
  fmt.Println(p.firstName + " " + p.lastName)
  fmt.Println((*p).firstName + " " + (*p).lastName)
}

func main() {
  p := person{
    firstName: "James",
    lastName:  "Bond",
  }
  name(&p)
}
```

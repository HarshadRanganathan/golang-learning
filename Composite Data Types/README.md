[Array](#array)

[Slice](#slice)

- [Slice expressions](#slice-expressions)

- [Appending elements](#appending-elements)

- [Removing elements](#removing-elements)

- [Slice Iteration](#slice-iteration)

- [Slice Sort](#slice-sort)

- [Memory allocation](#memory-allocation)

[Map](#map)

- [Checking if a key is present](#checking-if-a-key-is-present)

- [Adding entries](#adding-entries)

- [Removing entries](#removing-entries)

- [Map Iteration](#map-iteration)

[Struct](#struct)

- [Struct in map](#struct-in-map)

- [Embedded fields](#embedded-fields)

- [Promoted fields](#promoted-fields)

- [Anonymous structures](#anonymous-structures)

- [Struct sort](#struct-sort)

## Array

An array is a numbered sequence of elements of a single type, called the element type.

```go
package main

import "fmt"

func main() {
  x := [5]int{1, 2, 3, 4, 5}

  for i, v := range x {
    fmt.Printf("%v %v\n", i, v)
  }
}
```

Output:

```text
0 1
1 2
2 3
3 4
4 5
```

## Slice

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.

Slices hold references to an underlying array, and if you assign one slice to another, both refer to the same array. If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, analogous to passing a pointer to the underlying array.

```go
package main

import "fmt"

func main() {
  x := []int{1, 2, 3, 4, 5}

  fmt.Printf("%v", cap(x)) // 5 - capacity of underlying array
}
```

### Slice expressions

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice.

```go
package main

import "fmt"

func main() {
  x := []int{1, 2, 3, 4, 5}

  fmt.Println(x[:3]) // [1 2 3]
  fmt.Println(x[3:]) // [4 5]
  fmt.Println(x[0:5]) // [1 2 3 4 5]
}
```

### Appending elements

The variadic function `append` appends zero or more values x to s of type S, which must be a slice type, and returns the resulting slice, also of type S.

```go
package main

import "fmt"

func main() {
  x := []int{1, 2, 3, 4, 5}
  y := []int{6, 7, 8, 9, 10}

  z := append(x, y...) // y followed by ... appends the values of y to x of type int

  fmt.Println(z) // [1 2 3 4 5 6 7 8 9 10]
}
```

### Removing elements

In order to remove elements from a slice and preserve the order, make use of the built in `append` function.

```go
package main

import "fmt"

func main() {
  x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  y := append(x[10:], x[5:]...) // new slice containing the last 5 values

  fmt.Println(y) // [6 7 8 9 10]
}
```

### Slice Iteration

Iterate the slice using the `range clause` in a for loop.

```go
package main

import (
  "fmt"
)

func main() {
  x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

  for i, v := range x {
    fmt.Println(i, v)
  }
}
```

### Slice Sort

Package sort provides primitives for sorting slices and user-defined collections.

``Ints`` sorts a slice of ints in increasing order.

``Strings`` sorts a slice of strings in increasing order.

```go
package main

import (
  "fmt"
  "sort"
)

func main() {
  num := []int{45, 32, 80, 90, 34, 12, 02}
  sort.Ints(num)
  fmt.Println(num)

  s := []string{"Money Penny", "James Bond"}
  sort.Strings(s)
  fmt.Println(s)
}
```

Output:
```text
[2 12 32 34 45 80 90]
[James Bond Money Penny]
```

### Memory allocation

You can make a slice using the built-in `make` function which takes in the length and capacity of the slice.

```go
package main

import "fmt"

func main() {
  x := make([]int, 10) // makes a slice of length and capacity 10

  fmt.Println(len(x))
  fmt.Println(cap(x))
  fmt.Println(x)

  // copy can handle source and destination slices that share the same underlying array

  copy(x, []int{1, 2, 3, 4, 5}) // copies the data to the destination slice
  fmt.Println(x)

  copy(x[5:10], []int{6, 7, 8, 9, 10}) // copies the data to the destination slice utilizing the underlying capacity
  fmt.Println(x)
}
```

Output:

```text
10
10
[0 0 0 0 0 0 0 0 0 0]
[1 2 3 4 5 0 0 0 0 0]
[1 2 3 4 5 6 7 8 9 10]
```

## Map

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.

```go
package main

import "fmt"

func main() {
  // construct a map with key type as string and element type as slice of string ([] string)
  x := map[string][]string{
    "citrus":   {"oranges", "limes"},
    "tropical": {"bananas", "mangoes"},
  }

  fmt.Println(x) // map[citrus:[oranges limes] tropical:[bananas mangoes]]

  fmt.Println(x["citrus"]) // [oranges limes]
}
```

### Checking if a key is present

To check if the map contains a key, make use of `Index expressions`.

```go
package main

import "fmt"

func main() {
  x := map[string][]string{
    "citrus":   {"oranges", "limes"},
    "tropical": {"bananas", "mangoes"},
  }

  v, ok := x["citrus"] // index expression yields the value of map and sets the value of 'ok' to true if the key is present; false otherwise
  fmt.Println(v, ok)   // [oranges limes] true

  // use of _ discards the value
  if _, ok := x["melons"]; !ok {
    fmt.Println("Melons are missing!") // prints if 'melons' key is missing in the map
  }
}
```

### Adding entries

You can add new entries to the map by using `index expressions`.

```go
package main

import "fmt"

func main() {
  x := map[string][]string{
    "citrus":   {"oranges", "limes"},
    "tropical": {"bananas", "mangoes"},
  }

  x["melons"] = []string{"watermelons", "honeydew"} // assign slice of values to key 'melons' in the map

  fmt.Println(x) // map[citrus:[oranges limes] melons:[watermelons honeydew] tropical:[bananas mangoes]]
}
```

### Removing entries

Use the built in function `delete` to remove contents from the map.

```go
package main

import "fmt"

func main() {
  x := map[string][]string{
    "citrus":   {"oranges", "limes"},
    "tropical": {"bananas", "mangoes"},
  }

  delete(x, "tropical")

  fmt.Println(x) // map[citrus:[oranges limes]]
}

```

### Map Iteration

Use `range` clause to iterate over the map.

```go
package main

import "fmt"

func main() {
  x := map[string][]string{
    "citrus":   {"oranges", "limes"},
    "tropical": {"bananas", "mangoes"},
  }

  for k, v := range x {
    fmt.Println(k, v)
  }
}
```

Output:

```text
citrus [oranges limes]
tropical [bananas mangoes]
```

## Struct

A struct is a sequence of named elements, called fields, each of which has a name and a type.

Within a struct, non-blank field names must be unique.

```go
package main

import "fmt"

type person struct {
  first      string
  last       string
  favFlavors []string
}

func main() {
  p := person{
    first:      "James",
    last:       "Bond",
    favFlavors: []string{"chocolate", "martini"},
  }

  fmt.Println(p.first)
  fmt.Println(p.last)

  for i, v := range p.favFlavors {
    fmt.Println(i, v)
  }
}
```

Output:

```text
James
Bond
0 chocolate
1 martini
```

### Struct in map

You can add struct elements to a map as shown below.

```go
package main

import (
  "fmt"
  "strings"
)

type person struct {
  first      string
  last       string
  favFlavors []string
}

func main() {
  p1 := person{
    first:      "James",
    last:       "Bond",
    favFlavors: []string{"chocolate", "martini"},
  }

  p2 := person{
    first:      "Money",
    last:       "Penney",
    favFlavors: []string{"strawberry", "vodka"},
  }

  m := map[string]person{
    strings.Join([]string{p1.first, p1.last}, " "): p1,
    strings.Join([]string{p2.first, p2.last}, " "): p2,
  }

  for k, v := range m {
    fmt.Println(k, v)
  }
}
```

Output:

```text
James Bond {James Bond [chocolate martini]}
Money Penney {Money Penney [strawberry vodka]}
```

### Embedded fields

A field declared with a type but no explicit field name is called an `embedded field`. An embedded field must be specified as a type name T.

```go
package main

import "fmt"

type vehicle struct {
  doors int
  color string
}

type truck struct {
  vehicle // embedded struct
  fourWheel bool
}

type sedan struct {
  vehicle // embedded struct
  luxury  bool
}

func main() {
  t := truck{
    vehicle: vehicle{
      doors: 2,
      color: "white",
    },
    fourWheel: true,
  }

  s := sedan{
    vehicle: vehicle{
      doors: 4,
      color: "silver",
    },
    luxury: true,
  }

  fmt.Println(t)
  fmt.Println(s)
}
```

Output:

```text
{{2 white} true}
{{4 silver} true}
```

### Promoted fields

A field or method f of an embedded field in a struct x is called `promoted` if x.f is a legal selector that denotes that field or method f.

You can access promoted fields either using the legal selector (or) using the field name.

```go
package main

import "fmt"

type vehicle struct {
  doors int
  color string
}

type truck struct {
  vehicle   // embedded struct with promoted fields doors & color
  fourWheel bool
}

func main() {
  t := truck{
    vehicle: vehicle{
      doors: 2,
      color: "white",
    },
    fourWheel: true,
  }

  fmt.Println(t.vehicle.color)
  fmt.Println(t.color) // color is a promoted field hence can be accessed without the selector
}
```

Output:

```text
white
white
```

### Anonymous structures

It is possible to declare structures without declaring a new type and these type of structures are called `anonymous structures`.

```go
package main

import "fmt"

func main() {
  s := struct {
    first     string
    friends   map[string]int
    favDrinks []string
  }{
    first: "James",
    friends: map[string]int{
      "Jack": 12,
      "Ryan": 14,
    },
    favDrinks: []string{"martini", "beer"},
  }

  fmt.Println(s)
  fmt.Println(s.first)
  fmt.Println(s.friends)
  fmt.Println(s.favDrinks)
}
```

Output:

```text
{James map[Jack:12 Ryan:14] [martini beer]}
James
map[Jack:12 Ryan:14]
[martini beer]
```

### Struct sort

You can sort a list of collections by using `Sort` function. It expects a type that satisfies [sort.Interface](https://golang.org/pkg/sort/#Interface).

```go
package main

import (
  "fmt"
  "sort"
)

type person struct {
  firstName string
  lastName  string
  age       int
}

// implements sort.Interface
type sortByAge []person

func (a sortByAge) Len() int           { return len(a) }
func (a sortByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
  p1 := person{firstName: "Kody", lastName: "Hunt", age: 34}
  p2 := person{firstName: "Kirsten", lastName: "Blackwell", age: 22}
  p3 := person{firstName: "Bernard", lastName: "Wheatley", age: 45}

  persons := []person{p1, p2, p3}
  sort.Sort(sortByAge(persons))

  fmt.Println(persons)
}
```

Output:

```text
[{Kirsten Blackwell 22} {Kody Hunt 34} {Bernard Wheatley 45}]
```
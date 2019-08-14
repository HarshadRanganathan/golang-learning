[Array](#array)

[Slice](#slice)

[Map](#map)

[Struct](#struct)

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

## Struct

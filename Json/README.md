[Json](#json)

- [Marshal](#marshal)

  - [Struct](#struct)

  - [Customize Encoding](#customize-encoding)

- [Unmarshal](#unmarshal)

## Json

import "encoding/json"

Package json implements encoding and decoding of JSON as defined in RFC 7159. 

### Marshal

Marshal returns the JSON encoding of value.

#### Struct

If you want to marshal a struct, the fields need to be `exported`.  Each exported struct field becomes a member of the object, using the field name as the object key.

```go
package main

import (
  "encoding/json"
  "fmt"
)

type person struct {
  // An identifier is exported if the first character of the identifier's name is a Unicode upper case letter
  FirstName string
  LastName  string
}

func main() {
  p1 := person{
    FirstName: "James",
    LastName:  "Bond",
  }

  p2 := person{
    FirstName: "Money",
    LastName:  "Penny",
  }

  persons := []person{p1, p2}
  fmt.Println(persons)

  data, err := json.Marshal(persons)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(data))
}
```

Output:

```text
[{James Bond} {Money Penny}]
[{"FirstName":"James","LastName":"Bond"},{"FirstName":"Money","LastName":"Penny"}]
```

#### Customize Encoding

The encoding of each struct field can be customized by the format string stored under the "json" key in the struct field's tag. The format string gives the name of the field, possibly followed by a comma-separated list of options.

The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value, defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.

```go
package main

import (
  "encoding/json"
  "fmt"
)

type person struct {
  FirstName string `json:"firstName,omitempty"`
  LastName  string `json:"lastName,omitempty"`
}

func main() {
  p1 := person{
    FirstName: "James",
    LastName:  "Bond",
  }

  p2 := person{
    FirstName: "Money",
    LastName:  "",
  }

  persons := []person{p1, p2}
  fmt.Println(persons)

  data, err := json.Marshal(persons)

  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(string(data))
}
```

Output:

```text
[{James Bond} {Money }]
[{"firstName":"James","lastName":"Bond"},{"firstName":"Money"}]
```

### Unmarshal

Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.

To unmarshal JSON into a pointer, Unmarshal first handles the case of the JSON being the JSON literal null. In that case, Unmarshal sets the pointer to nil. Otherwise, Unmarshal unmarshals the JSON into the value pointed at by the pointer. 

```go
package main

import (
  "encoding/json"
  "fmt"
)

type person struct {
  FirstName string `json:"firstName,omitempty"`
  LastName  string `json:"lastName,omitempty"`
}

func main() {
  s := `[{"FirstName":"James","LastName":"Bond"},{"FirstName":"Money","LastName":"Penny"}]`
  bs := []byte(s)

  var persons []person

  err := json.Unmarshal(bs, &persons)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Println(persons)
}
```
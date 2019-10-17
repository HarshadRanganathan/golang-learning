[Bcrypt](#bcrypt)

## Bcrypt

Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.

``GenerateFromPassword`` returns the bcrypt hash of the password at the given cost.

``CompareHashAndPassword`` compares a bcrypt hashed password with its possible plaintext equivalent. Returns nil on success, or an error on failure.

This is part of experimental package. Download the source code of the package using below command for the import to resolve.

```bash
go get golang.org/x/crypto/bcrypt
```

```go
package main

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
)

func main() {
  password := "admin"

  hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
  if err != nil {
    fmt.Println(err)
  }

  err = bcrypt.CompareHashAndPassword(hash, []byte(password))
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Password matches")
  }

  err = bcrypt.CompareHashAndPassword(hash, []byte("jedi"))
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Password matches")
  }
}
```

Output:

```text
Password matches
crypto/bcrypt: hashedPassword is not the hash of the given password
```
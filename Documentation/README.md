[Documentation](#documentation)

- [go doc](#go-doc)

- [godoc](#godoc)

## Documentation

To document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line.

Comments on package declarations should provide general package documentation.

```go
// Package hello provides function for division operation
package hello

// Divide takes in two operands and returns a result
func Divide(x int, y int) int {
	if y == 0 {
		panic("Divide by zero")
	}
	return x / y
}
```

### go doc

`go doc` shows the documentation for the package in the current directory.

```text
$ go doc
package hello // import "."

Package hello provides function for division operation

func Divide(x int, y int) int
```

`go doc <sym>[.<methodOrField>]` - If there is a symbol but no package, the package in the current directory is chosen.

```
$ go doc Divide
func Divide(x int, y int) int
    Divide takes in two operands and returns a result
```

### godoc

Godoc extracts and generates documentation for Go programs.

It runs as a web server and presents the documentation as a web page.

By default, godoc looks at the packages it finds via ``$GOROOT`` and ``$GOPATH`` (if set).

Incase, you want to generate godoc for your local package, make sure it has proper project structure eg: ``<GOPATH>/src/<package>/hello.go`` and run below command:

```bash
$ godoc -http=:6060
```

Open your browser and access your package docs at ``http://localhost:6060/pkg/<package_name>``

![godoc](godoc.png?raw=true)
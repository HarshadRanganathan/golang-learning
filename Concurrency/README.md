[Concurrency](#concurrency)

- [Goroutine](#goroutine)

- [WaitGroup](#waitgroup)

- [Mutex](#mutex)

- [Atomic](#atomic)

- [Channels](#channels)

  - [Range clause](#range-clause)

  - [Select statements](#select-statements)

- [Build package with race detection](#build-package-with-race-detection)

## Concurrency

### Goroutine

A goroutine is a lightweight thread managed by the Go runtime.

Goroutines run in the same address space, so access to shared memory must be synchronized.

```go
package main

import (
  "fmt"
  "runtime"
  "time"
)

func printEvnNos() {
  for i := 0; i <= 10; i++ {
    if i%2 == 0 {
      fmt.Println(i)
    }
    time.Sleep(100 * time.Millisecond)
  }
}

func printOddNos() {
  for i := 0; i <= 10; i++ {
    if i%2 != 0 {
      fmt.Println(i)
    }
    time.Sleep(100 * time.Millisecond)
  }
}

func main() {
  fmt.Println("CPUs: ", runtime.NumCPU())

  go printOddNos() // execution happens in a new goroutine
  printEvnNos()
}
```

### WaitGroup

A WaitGroup waits for a collection of goroutines to finish. 

The main goroutine calls Add to set the number of goroutines to wait for. 

Then each of the goroutines runs and calls Done when finished. 

At the same time, Wait can be used to block until all goroutines have finished.

```go
package main

import (
  "fmt"
  "runtime"
  "sync"
  "time"
)

var wg sync.WaitGroup

func printEvnNos() {
  for i := 0; i <= 10; i++ {
    if i%2 == 0 {
      fmt.Println(i)
    }
    time.Sleep(100 * time.Millisecond)
  }
}

func printOddNos() {
  for i := 0; i <= 10; i++ {
    if i%2 != 0 {
      fmt.Println(i)
    }
    time.Sleep(100 * time.Millisecond)
  }
  wg.Done() // decrements the WaitGroup counter by '1'
}

func main() {
  fmt.Println("CPUs: ", runtime.NumCPU())
  fmt.Println("GoRoutine: ", runtime.NumGoroutine())

  wg.Add(1) // wait for '1' goroutine to finish

  go printOddNos()
  printEvnNos()

  fmt.Println("GoRoutine: ", runtime.NumGoroutine())
  wg.Wait() // wait until all the goroutines have finsihed
}
```

Output:

```text
CPUs:  12
GoRoutine:  1
0
1
2
3
4
5
6
7
8
9
10
GoRoutine:  2
```

### Mutex

A Mutex is a mutual exclusion lock which will synchronize access to state.

`Lock` - Locks the state. If the lock is already in use, the calling goroutine blocks until the mutex is available.

`Unloack` - Unlocks the mutex.

```go
package main

import (
  "fmt"
  "sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func race() {
  mu.Lock()
  counter++
  mu.Unlock()
  wg.Done()
}

func main() {
  wg.Add(2)

  go race()
  go race()

  wg.Wait()

  fmt.Println(counter)
}
```

Output:

`--race` helps to determine if there are any race conditions

```bash
$ go run --race hello.go
2
```

In the absence of mutex usage, below will be the output:

```bash
$ go run --race hello.go
==================
WARNING: DATA RACE
Read at 0x000000606318 by goroutine 7:
  main.race()
      C:/GoWorkspace/test/hello.go:12 +0x45

Previous write at 0x000000606318 by goroutine 6:
  main.race()
      C:/GoWorkspace/test/hello.go:12 +0x61

Goroutine 7 (running) created at:
  main.main()
      C:/GoWorkspace/test/hello.go:20 +0x7e

Goroutine 6 (finished) created at:
  main.main()
      C:/GoWorkspace/test/hello.go:19 +0x66
==================
2
Found 1 data race(s)
exit status 66
```

### Atomic

Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

`AddInt64` - atomically adds delta to *addr and returns the new value.

```go
package main

import (
  "fmt"
  "sync"
  "sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func race() {
  atomic.AddInt64(&counter, 1)
  wg.Done()
}

func main() {
  wg.Add(2)

  go race()
  go race()

  wg.Wait()

  fmt.Println(counter)
}
```

Output:

```text
$ go run --race hello.go
2
```

### Channels

A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type.

```go
package main

import "fmt"

func main() {
  c := make(chan int)
  go func() { c <- 42 }() // send to channel
  fmt.Println(<-c) // receive from channel
}
```

The optional <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bidirectional. 

```go
package main

import "fmt"

func main() {
  cs := make(chan<- int) // send only channel
  go func() { cs <- 42 }()
  fmt.Println(<-cs) // error - receive from send only type

  cr := make(<-chan int)   // receive only channel
  go func() { cr <- 42 }() // error - send to receive only type
  fmt.Println(<-cr)
}
```

Below is an example of how to use send and receive channels to perform concurrent operations.

```go
package main

import "fmt"

func send(c chan<- int) {
  c <- 42
}

func receive(c <-chan int) {
  fmt.Println(<-c)
}

func main() {
  c := make(chan int)
  go send(c)
  receive(c)
}
```

#### Range clause

```go
package main

import "fmt"

func main() {
  c := make(chan int)

  go func() {
    for i := 0; i < 5; i++ {
      c <- i
    }
    close(c) // channel needs to be closed to indicate no more values will be sent to it
  }()

  for v := range c {
    fmt.Println(v)
  }
}
```

Output:

```text
0
1
2
3
4
```

#### Select statements

A "select" statement chooses which of a set of possible send or receive operations will proceed. It looks similar to a "switch" statement but with the cases all referring to communication operations.

Result is a set of channels to receive from or send to, and the corresponding values to send.

```go
package main

import "fmt"

func publish(odd chan int, even chan int, quit chan int) {
  for i := 0; i < 10; i++ {
    if i%2 == 0 {
      even <- i
    } else {
      odd <- i
    }
  }
  close(quit)
}

func main() {
  odd := make(chan int)
  even := make(chan int)
  quit := make(chan int)

  go publish(odd, even, quit)

  for {
    select {
    case o := <-odd:
      fmt.Println("Received odd value: ", o)
    case e := <-even:
      fmt.Println("Received even value: ", e)
    case _, ok := <-quit:
      if !ok { // channel is closed
        close(odd)
        close(even)
        fmt.Println("Received quit")
      }
      return
    }
  }
}
```

Output:

```text
Received even value:  0
Received odd value:  1
Received even value:  2
Received odd value:  3
Received even value:  4
Received odd value:  5
Received even value:  6
Received odd value:  7
Received even value:  8
Received odd value:  9
Received quit
```

### Build package with race detection

Build compiles the packages named by the import paths, along with their dependencies.

When compiling a single main package, build writes the resulting executable to an output file.

`-race` flag enables race detection

```bash
go build --race
```

You can then run the resulting executable.

```bash
./test.exe
```
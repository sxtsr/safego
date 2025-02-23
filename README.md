# safego

SafeGo is a simple Go package that provides a basic safety template for goroutine function execution with panic recovery and context support.

## Installation

```bash
go get github.com/sxtsr/safego
```

## Features
- Panic recovery in goroutines
- Context-aware goroutine execution
- Simple and lightweight implementation

## Usage Examples

### Basic Usage
```go
package main

import (
    "fmt"
    "github.com/sxtsr/safego"
)

func main() {
    // Example of safe goroutine execution
    safego.Go(func() {
        fmt.Println("Executing safely in goroutine")
        
        // Even if panic occurs, it will be recovered and logged
        panic("test panic")
    })
}
```

### With Context
```go
package main

import (
    "context"
    "fmt"
    "github.com/sxtsr/safego"
    "time"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    safego.GoContext(ctx, func() {
        fmt.Println("Executing with context")
        // Will not execute if context is canceled
    })
}
```

### Multiple Goroutines
```go
package main

import (
    "fmt"
    "github.com/sxtsr/safego"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        i := i
        
        safego.Go(func() {
            defer wg.Done()
            fmt.Printf("Processing item %d\n", i)
        })
    }
    
    wg.Wait()
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

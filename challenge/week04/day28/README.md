# Day 27 of #66DaysOfGo

_Last update:  Aug 16, 2023_.

---

Today, I've continued with the Design Patterns series, with the Facade.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Facade pattern

This pattern simplify complex systems. Similar to architectural facades, it offers a simplified interface that hides intricate underlying code. It enhances readability, provides context-specific interfaces, and aids in refactoring monolithic systems. This pattern involves a wrapper class that presents a clear interface while masking complexities.

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/facade/structure-indexed-2x.png" alt="Facade Pattern UML example" width="450"/>

### Code example

```go
package main

import "fmt"

type ComplexETL struct{}

func (etl *ComplexETL) Extract() {
    fmt.Println("Extracting...")
}
func (etl *ComplexETL) Transform() {
    fmt.Println("Transforming...")
}
func (etl *ComplexETL) Load() {
    fmt.Println("Loading...")
}

type ETLFacade struct {
    ETL *ComplexETL
}

func (etl *ETLFacade) Process() {
    etl.ETL.Extract()
    etl.ETL.Transform()
    etl.ETL.Load()
}

func main() {
    etl := &ETLFacade{&ComplexETL{}}
    etl.Process()
}

```

```bash
$ go run facade.go
Extracting...
Transforming...
Loading...
```

---

## References

- []()

# Day 25 of #66DaysOfGo

_Last update:  Aug 13, 2023_.

---

Today, I've continued with the Design Patterns series, with the Bridge.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Bridge pattern

The bridge pattern approach is to separate abstraction from implementation, allowing independent variation. It uses encapsulation, aggregation, and possibly inheritance to split responsibilities into different classes. It's useful for frequent class and functionality changes, it's like two layers of abstraction (not to be confused with the adapter pattern, where implementation can be deferred until abstraction is used).

### UML diagram

<img src="https://refactoring.guru/images/patterns/diagrams/bridge/structure-en-2x.png" alt="Bridge Pattern UML example" width="550"/>

### Code example

```go
package main

import "fmt"

type CloudComputeService interface {
    runProcess()
    setArchitecture(Architecture)
}

type VM struct {
    architecture Architecture
}

type Serverless struct {
    architecture Architecture
}

func (v *VM) runProcess() {
    fmt.Print("Running in a Virtual Machine Compute service")
    v.architecture.runProcess()
}
func (v *VM) setArchitecture(architecture Architecture) {
    v.architecture = architecture
}

func (s *Serverless) runProcess() {
    fmt.Print("Running in a Serverless Compute service")
    s.architecture.runProcess()
}
func (s *Serverless) setArchitecture(architecture Architecture) {
    s.architecture = architecture
}

type Architecture interface {
    runProcess()
}

type IntelArchitecture struct{}
type AmdArchitecture struct{}

func (i *IntelArchitecture) runProcess() {
    fmt.Print(" (running process in Intel architecture...)")
}
func (a *AmdArchitecture) runProcess() {
    fmt.Print(" (running process in AMD architecture...)")
}

func main() {
    intelArchitecture := IntelArchitecture{}
    vm := VM{
        architecture: &intelArchitecture,
    }
    vm.runProcess()

    fmt.Println()

    amdArchitecture := AmdArchitecture{}
    serverless := Serverless{
        architecture: &amdArchitecture,
    }
    serverless.runProcess()
}
```

```bash
$ go run bridge.go
Running in a Virtual Machine Compute service (running process in Intel architecture...)
Running in a Serverless Compute service (running process in AMD architecture...)
```

---

## References

- [https://refactoring.guru/design-patterns/bridge](https://refactoring.guru/design-patterns/bridge)

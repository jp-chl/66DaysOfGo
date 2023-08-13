# Day 23 of #66DaysOfGo

_Last update:  Aug 11, 2023_.

---

Today, I've continued with the Design Patterns series, with the Prototype.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Prototype pattern

The Prototype pattern enables object cloning by delegating the process to the objects themselves, using a shared interface. This allows cloning without tightly coupling code to object classes. Objects implementing this interface contain a clone method, creating new instances and transferring old object's field values. Cloning is beneficial for objects with many fields or configurations, offering an alternative to subclassing.

### UML diagram

<img src="https://i1.wp.com/golangbyexample.com/wp-content/uploads/2019/10/Prototype-Pattern.jpg?w=639&ssl=1" alt="Prototype Pattern UML example" width="450"/>

### Code example

```go
package main

import "fmt"

type Car interface {
    print()
    clone() Car
}

type Bmw struct {
    year  int
    color string
}

func (b *Bmw) clone() Car {
    return &Bmw{
        year:  b.year,
        color: b.color,
    }
}

func (b *Bmw) print() {
    fmt.Printf("Brand: BMW, Year: %d, Color: %s\n", b.year, b.color)
}

type Porsche struct {
    brand string
    year  int
    power int
}

func (p *Porsche) clone() Car {
    return &Porsche{
        year:  p.year,
        power: p.power,
    }
}

func (p *Porsche) print() {
    fmt.Printf("Brand: Porsche, Year: %d, Power: %d\n", p.year, p.power)
}

func main() {
    car1 := &Bmw{
        year:  2000,
        color: "red",
    }
    car2 := &Porsche{
        year:  2000,
        power: 500,
    }
    demoCar(car1)
    demoCar(car2)

    fmt.Println("\nCloning...")
    car3 := car1.clone()
    car4 := car2.clone()
    demoCar(car3)
    demoCar(car4)
}

func demoCar(c Car) {
    fmt.Printf("Variable memory address: %v. ", &c)
    c.print()
}
```

```bash
$ go run prototype.go
Variable memory address: 0xc000096230. Brand: BMW, Year: 2000, Color: red
Variable memory address: 0xc000096250. Brand: Porsche, Year: 2000, Power: 500

Cloning...
Variable memory address: 0xc000096260. Brand: BMW, Year: 2000, Color: red
Variable memory address: 0xc000096280. Brand: Porsche, Year: 2000, Power: 500
```

---

## References

- [https://refactoring.guru/design-patterns/prototype](https://refactoring.guru/design-patterns/prototype)

# Day 32 of #66DaysOfGo

_Last update:  Aug 20, 2023_.

---

Today, I've continued with the Design Patterns series, with the Command.

---

## Versions used

- macOS Monterrey 12.2
- go: 1.20.6

---

## Command pattern

This pattern involves encapsulating information about an action or event in an object for later execution. It uses command, receiver, invoker, and client components. Command objects store method details and receiver, while invoker triggers execution. This pattern reduces coupling and enables flexible configuration of requests through separate objects. It aligns with first-class and higher-order functions in functional programming, offering flexibility and reusability in software design.

### UML diagram

<img src="https://upload.wikimedia.org/wikipedia/commons/c/c8/W3sDesign_Command_Design_Pattern_UML.jpg" alt="Command Pattern UML example" width="750"/>

### Code example

```go
package main

import "fmt"

type Command interface {
    Execute()
}

// Receiver
type Light struct {
    isOn bool
}

func (l *Light) TurnOn() {
    l.isOn = true
    fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
    l.isOn = false
    fmt.Println("Light is off")
}

// Concrete Command
type TurnOnCommand struct {
    light *Light
}

func NewTurnOnCommand(light *Light) *TurnOnCommand {
    return &TurnOnCommand{light: light}
}

func (c *TurnOnCommand) Execute() {
    c.light.TurnOn()
}

// Concrete Command
type TurnOffCommand struct {
    light *Light
}

func NewTurnOffCommand(light *Light) *TurnOffCommand {
    return &TurnOffCommand{light: light}
}

func (c *TurnOffCommand) Execute() {
    c.light.TurnOff()
}

// Invoker
type RemoteControl struct {
    command Command
}

func (rc *RemoteControl) SetCommand(command Command) {
    rc.command = command
}

func (rc *RemoteControl) PressButton() {
    rc.command.Execute()
}

func main() {
    light := &Light{}
    turnOnCommand := NewTurnOnCommand(light)
    turnOffCommand := NewTurnOffCommand(light)

    remoteControl := &RemoteControl{}

    remoteControl.SetCommand(turnOnCommand)
    remoteControl.PressButton()

    remoteControl.SetCommand(turnOffCommand)
    remoteControl.PressButton()
}

```

```bash
$ go run command.go
Light is on
Light is off
```

---

## References

- [https://refactoring.guru/design-patterns/command](https://refactoring.guru/design-patterns/command)

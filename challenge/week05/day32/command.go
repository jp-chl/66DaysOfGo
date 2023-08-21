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

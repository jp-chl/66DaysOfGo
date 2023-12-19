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

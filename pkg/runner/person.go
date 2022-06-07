package runner

import "fmt"

type PersonProcess struct {
	Name string
}

func NewPersonProcess(name string) *PersonProcess { return &PersonProcess{Name: name} }

func (u PersonProcess) GetName() string { return u.Name }

func (u PersonProcess) Execute() {
	fmt.Println("PersonProcess")
}

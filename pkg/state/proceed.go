package state

import (
	"fmt"
)

type Proceed struct {
	Job *Job
}

func (p *Proceed) Do(defers chan bool, processName string) {
	fmt.Println("proceed", processName)
	defers <- false
}

func (p *Proceed) Prompt(processName string) {
}

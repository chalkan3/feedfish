package state

import (
	"fmt"
	"time"
)

type Process struct {
	Job    *Job
	Script string
}

func (p *Process) Do(defers chan bool, processName string) {
	Shell(p.Script)
	p.Job.SetState(p.Job.Proceed)
	defers <- true

}

func (p *Process) Prompt(processName string) {
	fmt.Println("Processing", processName)
	time.Sleep(time.Second * 3)

}

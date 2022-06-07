package state

import (
	"fmt"
	"time"
)

type Process struct {
	Job *Job
}

func (p *Process) Do(defers chan bool, processName string) {
	fmt.Println("processing", processName)
	time.Sleep(time.Second * 15)
	fmt.Println("end processing")
	p.Job.SetState(p.Job.Proceed)
	defers <- true

}

func (p *Process) Prompt(processName string) {
	fmt.Println("Processing", processName)
	time.Sleep(time.Second * 3)

}

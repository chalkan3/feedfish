package runner

import (
	state "t/pkg/state"
)

type Process struct {
	Name string
	Job  *state.Job
}

func NewProcess(name string, initialScript string, processingScript string, proceedScript string) *Process {
	return &Process{Name: name, Job: state.NewJob(initialScript, processingScript, proceedScript)}
}

func (u Process) GetName() string { return u.Name }

func (u Process) Execute() {
	u.Job.Do(u.Name)
}

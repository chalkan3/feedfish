package state

import (
	"fmt"
	"os/exec"
)

type State interface {
	Do(defers chan bool, processName string)
	Prompt(processName string)
}

type Job struct {
	Initial      State
	Processing   State
	Proceed      State
	CurrentState State
	defers       chan bool
}

func NewJob(initialScript, processingScript, proceedScript string) *Job {
	job := &Job{}
	ProceedState := &Proceed{Job: job, Script: proceedScript}
	ProcessingState := &Process{Job: job, Script: processingScript}
	InitialState := &Initial{Job: job, Script: initialScript}

	job.SetState(InitialState)
	job.Initial = InitialState
	job.Proceed = ProceedState
	job.Processing = ProcessingState
	job.defers = make(chan bool)

	return job
}

func (j *Job) SetState(s State) *Job {
	j.CurrentState = s
	return j
}

func (j *Job) Do(processName string) error {
	go j.CurrentState.Do(j.defers, processName)

	for {
		select {
		case defers := <-j.defers:
			if !defers {
				close(j.defers)
				return nil
			}
			go j.CurrentState.Do(j.defers, processName)
		default:
			j.CurrentState.Prompt(processName)
		}
	}

}

func Shell(script string) {
	cmd := exec.Command("/bin/sh", "-c", script)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(stdout))
}

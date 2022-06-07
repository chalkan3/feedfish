package state

import (
	"fmt"
	"time"
)

type Initial struct {
	Job *Job
}

func (i *Initial) Do(defers chan bool, processName string) {
	fmt.Println("initial", processName)
	time.Sleep(time.Second * 15)
	i.Job.SetState(i.Job.Processing)
	defers <- true
}

func (i *Initial) Prompt(processName string) {
	fmt.Println("Initializing process", processName)
	time.Sleep(time.Second * 3)

}

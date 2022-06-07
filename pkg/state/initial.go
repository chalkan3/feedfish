package state

import (
	"fmt"
	"time"
)

type Initial struct {
	Job    *Job
	Script string
}

func (i *Initial) Do(defers chan bool, processName string) {
	Shell(i.Script)
	i.Job.SetState(i.Job.Processing)
	fmt.Println("=================================================================================")
	defers <- true
}

func (i *Initial) Prompt(processName string) {
	fmt.Println(fmt.Sprintf("=======================Initializing process[%s]=============================", processName), "\n")
	time.Sleep(time.Second * 3)

}

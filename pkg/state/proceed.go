package state

type Proceed struct {
	Job    *Job
	Script string
}

func (p *Proceed) Do(defers chan bool, processName string) {
	Shell(p.Script)
	defers <- false
}

func (p *Proceed) Prompt(processName string) {
}

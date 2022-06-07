package state

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

func NewJob() *Job {
	job := &Job{}
	ProceedState := &Proceed{Job: job}
	ProcessingState := &Process{Job: job}
	InitialState := &Initial{Job: job}

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
				return nil
			}
			go j.CurrentState.Do(j.defers, processName)
		default:
			j.CurrentState.Prompt(processName)
		}
	}

}

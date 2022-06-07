package runner

type Runner struct {
	transversal *Stack
	Status      *Status
	Recovery    *ProcessHistory
}

func NewRunner(root *Node) *Runner {
	ui := &Runner{transversal: new(Stack), Status: new(Status), Recovery: &ProcessHistory{status: make([]*StatusMemento, 0)}}
	ui.moveLeft(root)

	return ui
}

func (u *Runner) moveLeft(current *Node) {
	for current != nil {
		u.transversal.Push(current)
		current = current.Left
	}
}

func (u *Runner) Next() process {
	if !u.HasNext() {
		return nil
	}

	current := u.transversal.Pop()
	u.Status.setState(current)
	u.Recovery.AddMemento(u.Status.createMemento())
	if current.Right != nil {
		u.moveLeft(current.Right)
	}

	return current.Process
}

func (u *Runner) HasNext() bool {
	return len(u.transversal.items) > 0
}

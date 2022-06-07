package runner

type Status struct {
	state *Node
}

type StatusMemento struct {
	state *Node
}

func (m *StatusMemento) getSavedState() *Node {
	return m.state
}

func (e *Status) createMemento() *StatusMemento {
	return &StatusMemento{state: e.state}
}

func (e *Status) restoreMemento(m *StatusMemento) {
	e.state = m.getSavedState()
}

func (e *Status) setState(node *Node) {
	e.state = node
}

func (e *Status) getState() *Node {
	return e.state
}

type ProcessHistory struct {
	status []*StatusMemento
}

func (c *ProcessHistory) AddMemento(m *StatusMemento) {
	c.status = append(c.status, m)
}

func (c *ProcessHistory) getMemento(index int) *StatusMemento {
	return c.status[index]
}

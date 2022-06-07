package runner

type Stack struct {
	items []*Node
}

func (s *Stack) Push(i *Node) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() *Node {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

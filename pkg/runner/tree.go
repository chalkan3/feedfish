package runner

type process interface {
	Execute()
}

type ProcessTree struct {
	head *Node
}

func NewProcessTree() *ProcessTree {
	return &ProcessTree{
		head: &Node{Key: 0, Process: NewProcess("teste")},
	}
}

func (pt *ProcessTree) Insert(k int, p process) *ProcessTree {
	pt.head.Insert(k, p)
	return pt
}

func (pt *ProcessTree) Search(k int) process {
	return pt.head.Search(k)
}

func (pt *ProcessTree) Root() *Node {
	return pt.head
}

type Node struct {
	Key     int
	Process process
	Left    *Node
	Right   *Node
}

func (n *Node) Insert(k int, p process) {
	node := &Node{Key: k, Process: p}

	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = node
		} else {
			n.Right.Insert(k, p)
		}
	} else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = node
		} else {
			n.Left.Insert(k, p)
		}
	}
}

func (n *Node) Search(k int) process {
	if n == nil {
		return nil
	}
	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return n.Process
}

// tree := NewProcessTree().
// 	Insert(100, &Person{Name: "igor"}).
// 	Insert(2, &Person{Name: "guica"}).
// 	Insert(40, &Person{Name: "lady"})

package simulator

// Node struct contains all the node state.
type Node struct {
	validators int
}

// NewNode returns a fully initialized node.
func NewNode() *Node {
	node := new(Node)
	node.reset()

	return node
}

func (node *Node) reset() {
	node.validators = 0
}

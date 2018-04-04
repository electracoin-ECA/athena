package simulator

const maxSimulatorNodes = 10

// Controller struct contains all the controller state.
type Controller struct {
	nodes []*Node
}

// NewController returns a fully initialized controller.
func NewController() (c *Controller) {
	c = new(Controller)
	c.reset()
	return
}

func (c *Controller) nodesLength() int {
	return len(c.nodes)
}

func (c *Controller) incrementNodes() {
	c.nodes = append(c.nodes, NewNode())
}

func (c *Controller) decrementNodes() {
	l := c.nodesLength()
	if l == 0 {
		return
	}

	c.nodes = c.nodes[0 : l-1]
}

func (c *Controller) reset() {
	c.nodes = make([]*Node, 0, 0)
}

package JSONExplorer

import jsonvalue "github.com/Andrew-M-C/go.jsonvalue"

type Container struct {
	inner  *jsonvalue.V
	child  []*Container
	islast bool
	isleaf bool
	value  string
}

func (c *Container) Draw() {
}

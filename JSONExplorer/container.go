package JSONExplorer

import jsonvalue "github.com/Andrew-M-C/go.jsonvalue"

type Container struct {
	inner    *jsonvalue.V
	Child    []*Container
	isfirst  bool
	islast   bool
	isbottom bool
	isleaf   bool
	key      string
	value    string
	level    int
}

func (c *Container) IsLeaf() bool {
	return c.isleaf
}

func (c *Container) Value() string {
	return c.value
}

func (c *Container) Level() int {
	return c.level
}

func (c *Container) IsLast() bool {
	return c.islast
}

func (c *Container) IsFirst() bool {
	return c.isfirst
}

func (c *Container) IsBottom() bool {
	return c.isbottom
}

func (c *Container) Key() string {
	return c.key
}

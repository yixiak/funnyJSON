package JSONExplorer

import jsonvalue "github.com/Andrew-M-C/go.jsonvalue"

// 迭代器设计模式
// 需要实现HasNext()和GetNext()方法，默认实现为中序遍历

type ContainerIter struct {
	inner  []*Container
	maxlen int
	index  int
}

func CreateIterator(json *jsonvalue.V) *ContainerIter {
	inner := make([]*Container, 0, 10)
	root := &Container{
		inner:    nil,
		isbottom: false,
		isleaf:   false,
		Child:    []*Container{},
		level:    0,
	}
	var dfs func(container *Container, jsonV *jsonvalue.V, isroot bool, isbottom bool, level int)

	dfs = func(container *Container, jsonV *jsonvalue.V, isroot bool, isbottom bool, level int) {
		childlen := jsonV.Len()
		index := 0
		for k, v := range jsonV.ForRangeObj() {
			index++
			if v.ValueType() == jsonvalue.Object {
				new_container := &Container{
					inner:    v,
					islast:   false,
					isleaf:   false,
					isbottom: false,
					key:      k,
					value:    "",
					Child:    []*Container{},
					level:    level + 1,
				}
				inner = append(inner, new_container)
				container.Child = append(container.Child, new_container)
				if isroot && index == childlen {
					if v.Len() == 0 {
						new_container.isbottom = true
					} else {
						dfs(new_container, v, false, true, level+1)
					}
				} else {
					dfs(new_container, v, false, isbottom, level+1)
				}

				if isbottom && index == childlen && v.Len() == 0 {
					new_container.isbottom = true
				}

				if index == childlen {
					new_container.islast = true
				}

				if isroot && index == 1 {
					new_container.isfirst = true
				}

			} else {
				new_leaf := &Container{
					inner:    v,
					isleaf:   true,
					islast:   false,
					isbottom: false,
					key:      k,
					value:    v.String(),
					level:    level + 1,
				}
				if isroot && index == childlen {
					new_leaf.isbottom = true
				}
				if isbottom && index == childlen {
					new_leaf.isbottom = true
				}
				if index == childlen {
					new_leaf.islast = true
				}
				if isroot && index == 1 {
					new_leaf.isfirst = true
				}
				container.Child = append(container.Child, new_leaf)
				inner = append(inner, new_leaf)
			}
		}
	}
	inner = append(inner, root)
	maxlen := int(1.7 * float32(getMaxlen(json, 1)))
	dfs(root, json, true, false, 0)

	return &ContainerIter{
		inner:  inner,
		maxlen: maxlen,
	}

}

func getMaxlen(V *jsonvalue.V, depth int) int {
	prex := depth * 3
	maxlen := prex
	if V != nil {
		for k, v := range V.ForRangeObj() {
			thislen := prex
			prex += len(k)
			//fmt.Printf("i: %v, v: %v\n", i, v)
			if v.ValueType() == jsonvalue.String {
				thislen += len(v.String()) + 1
				if thislen > maxlen {
					maxlen = thislen
				}
			}
			if v.ValueType() != jsonvalue.Null {
				x := getMaxlen(v, depth+1)
				if x > maxlen {
					maxlen = x
				}
			}

		}
	}
	return maxlen
}

func (c *ContainerIter) HasNext() bool {
	return c.index < len(c.inner)
}

func (c *ContainerIter) GetNext() *Container {
	if c.HasNext() {
		c.index++
		return c.inner[c.index-1]

	}
	return nil
}

func (c *ContainerIter) GetMaxlen() int {
	return c.maxlen
}

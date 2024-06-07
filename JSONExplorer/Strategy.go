package JSONExplorer

import "fmt"

type TreeStrategy struct {
	style    StyleFamily
	leaficon LeafIcon
	nodeicon NodeIcon
}

func (s *TreeStrategy) Draw(iter Iterator, style StyleFamily, leaficon LeafIcon, nodeIcon NodeIcon) {
	s.style = style
	s.leaficon = leaficon
	s.nodeicon = nodeIcon
	root := iter.GetNext()
	childlen := len(root.Child)
	index := 0
	for iter.HasNext() {
		index += 1
		prex := []rune{}

		container := iter.GetNext()

		var output []rune

		if container.IsLast() {
			if container.IsBottom() {
				output = []rune(s.style.Get_prefix_end_g())
			} else {
				output = []rune(s.style.Get_prefix_endleaf())
			}
		} else if container.IsFirst() {
			output = []rune(s.style.Get_prefix_first())
		} else {
			output = []rune(s.style.Get_prefix())
		}
		if container.IsLeaf() {
			output = append(output, []rune(s.leaficon.GetLeaf_Icon())...)
		} else {
			output = append(output, []rune(s.nodeicon.GetNode_Icon())...)
		}
		output = append(output, []rune(container.Key())...)
		output = append(output, []rune(fmt.Sprintf(":%s", container.Value()))...)

		fmt.Println(string(output))

		for _, child := range container.Child {
			iter.GetNext()
			if index == childlen {
				prex = append(prex, []rune(s.style.Get_symbol())...)
				s._draw(iter, prex, child, true)
			} else {
				prex = append(prex, []rune(s.style.Get_symbol_last())...)
				s._draw(iter, prex, child, false)
			}
		}
	}
}

func (s *TreeStrategy) _draw(iter Iterator, prex []rune, container *Container, islast bool) {
	output := []rune{}
	output = append(output, prex...)

	if container.IsLast() {
		if container.IsBottom() {
			output = append(output, []rune(s.style.Get_prefix_end_g())...)
		} else {
			output = append(output, []rune(s.style.Get_prefix_endleaf())...)
		}
	} else if container.IsFirst() {
		output = append(output, []rune(s.style.Get_prefix_first())...)
	} else {
		output = append(output, []rune(s.style.Get_prefix())...)
	}
	if container.IsLeaf() {
		output = append(output, []rune(s.leaficon.GetLeaf_Icon())...)
	} else {
		output = append(output, []rune(s.nodeicon.GetNode_Icon())...)
	}
	output = append(output, []rune(container.Key())...)
	output = append(output, []rune(fmt.Sprintf(":%s", container.Value()))...)

	fmt.Println(string(output))

	for _, child := range container.Child {
		iter.GetNext()
		s._draw(iter, append(prex, []rune(s.style.Get_symbol())...), child, islast)
	}

}

type RecStrategy struct {
}

func (s *RecStrategy) Draw(iter Iterator, style StyleFamily, leaficon, LeafIcon, nodeIcon NodeIcon) {}

package JSONExplorer

import "fmt"

// 策略设计模式
// 只需要在此添加想要的策略，并实现Draw方法即可

type TreeStrategy struct {
	style    StyleFamily
	leaficon LeafIcon
	nodeicon NodeIcon
}

func (s *TreeStrategy) Draw(iter Iterator, style StyleFamily, leaficon LeafIcon, nodeIcon NodeIcon) {
	s.style = style
	s.leaficon = leaficon
	s.nodeicon = nodeIcon

	container := iter.GetNext()
	//childlen := len(container.Child)
	for _, child := range container.Child {
		iter.GetNext()
		s._draw(iter, []rune{}, child)
	}

}

func (s *TreeStrategy) _draw(iter Iterator, prex []rune, container *Container) {
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
		if container.IsLast() {
			s._draw(iter, append(prex, []rune(s.style.Get_symbol_last())...), child)
		} else {
			s._draw(iter, append(prex, []rune(s.style.Get_symbol())...), child)
		}
	}
}

type RecStrategy struct {
	style    StyleFamily
	leaficon LeafIcon
	nodeicon NodeIcon
}

func (s *RecStrategy) Draw(iter Iterator, style StyleFamily, leaficon LeafIcon, nodeIcon NodeIcon) {
	s.style = style
	s.leaficon = leaficon
	s.nodeicon = nodeIcon
	iter.GetNext()

	for iter.HasNext() {
		container := iter.GetNext()
		output := []rune{}

		symbol := []rune{}
		level := container.Level()
		for i := 1; i < level; i++ {
			if container.IsBottom() {
				if i == 1 {
					symbol = append(symbol, []rune(s.style.Get_symbol_left_last())...)
				} else {
					symbol = append(symbol, []rune(s.style.Get_symbol_last_mid())...)
				}
			} else {
				symbol = append(symbol, []rune(s.style.Get_symbol())...)
			}
		}
		output = append(output, symbol...)

		prefix := []rune{}

		if container.IsFirst() {
			prefix = append(prefix, []rune(s.style.Get_prefix_first())...)
		} else if container.IsBottom() {
			prefix = append(prefix, []rune(s.style.Get_prefix_end_g())...)
		} else {
			prefix = append(prefix, []rune(s.style.Get_prefix())...)
		}

		if container.IsLeaf() {
			prefix = append(prefix, []rune(s.leaficon.GetLeaf_Icon())...)
		} else {
			prefix = append(prefix, []rune(s.nodeicon.GetNode_Icon())...)
		}

		prefix = append(prefix, []rune(container.Key())...)
		prefix = append(prefix, []rune(fmt.Sprintf(":%s", container.Value()))...)
		output = append(output, prefix...)

		end := []rune{}
		for i := 0; i < iter.GetMaxlen()-len(output)-1; i++ {
			end = append(end, []rune(s.style.Get_end())...)
		}

		if container.IsFirst() {
			end = append(end, []rune(s.style.Get_end_first())...)
		} else if container.IsBottom() {
			end = append(end, []rune(s.style.Get_end_last())...)
		} else {
			end = append(end, []rune(s.style.Get_end_mid())...)
		}

		output = append(output, end...)
		fmt.Println(string(output))

	}
}

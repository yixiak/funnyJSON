package JSONExplorer

type TreeStyle struct {
}

func (tree *TreeStyle) Get_symbol_last() string {
	return "   "
	// spac*3
} // 最后一个子对象的前缀
func (tree *TreeStyle) Get_symbol() string {
	return "|  "
} // 正常子对象的前缀
func (tree *TreeStyle) Get_symbol_left_last() string {
	return " "
} // 左下角
func (tree *TreeStyle) Get_symbol_last_mid() string {
	return " "
}

// prefix
// 实际上指selfjson的前缀
func (tree *TreeStyle) Get_prefix_end_g() string {
	return "└─"
} // 全局的最后一个节点，一定在最底部
func (tree *TreeStyle) Get_prefix_endleaf() string {
	return "└─"
} //前json对象的最后一个节点，但下面还有其他的
func (tree *TreeStyle) Get_prefix() string {
	return "├─"
}
func (tree *TreeStyle) Get_prefix_first() string {
	return "├─"
}

// end
func (tree *TreeStyle) Get_end() string {
	return " "
} // 中间填充
func (tree *TreeStyle) Get_end_first() string {
	return " "
} // 右上角
func (tree *TreeStyle) Get_end_mid() string {
	return " "
} // 右边框
func (tree *TreeStyle) Get_end_last() string {
	return " "
} // 右下角

type TreeStyleFactory struct{}

func (factory *TreeStyleFactory) CreateStlyeFamily() StyleFamily {
	return &TreeStyle{}
}

type RecStyle struct{}

func (tree *RecStyle) Get_symbol_last() string {
	return "|  "

} // 最后一个子对象的前缀
func (tree *RecStyle) Get_symbol() string {
	return "|  "
} // 正常子对象的前缀
func (tree *RecStyle) Get_symbol_left_last() string {
	return "└──"
} // 左下角

func (tree *RecStyle) Get_symbol_last_mid() string {
	return "───"
}

// prefix
// 实际上指selfjson的前缀
func (tree *RecStyle) Get_prefix_end_g() string {
	return "└─"
} // 全局的最后一个节点，一定在最底部
func (tree *RecStyle) Get_prefix_endleaf() string {
	return "├─"
} //前json对象的最后一个节点，但下面还有其他的
func (tree *RecStyle) Get_prefix() string {
	return "├─"
}
func (tree *RecStyle) Get_prefix_first() string {
	return "┌─"
}

// end
func (tree *RecStyle) Get_end() string {
	return "─"
} // 中间填充
func (tree *RecStyle) Get_end_first() string {
	return "┐"
} // 右上角
func (tree *RecStyle) Get_end_mid() string {
	return "┤"
} // 右边框
func (tree *RecStyle) Get_end_last() string {
	return "┘"
} // 右下角

type RecStyleFactory struct{}

func (factory *RecStyleFactory) CreateStlyeFamily() StyleFamily {
	return &RecStyle{}
}

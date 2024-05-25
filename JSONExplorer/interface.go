package JSONExplorer

// 使用生成器模式，可以自定义配置
type Explorer interface {
	// ParseJSON
	// 根据文件名解析 json
	ParseJSON(jsonfile string) error

	// InitStyle
	// 选择对应的风格
	InitStyle(style string)

	// InitIcon
	// 选择对应的图标族
	InitIcon(IconFamily string)

	// Show
	// 最后输出
	Show()
}

type IconFamily interface {
	GetLeaf_prefixIcon() string
	GetLeaf_suffixIcon() string
	GetNode_prefixIcon() string
	GetNode_suffixIcon() string
	// Get_leafprefix() string
	// Get_leaftendconprefix() string
	// Get_leafend() string
	// Get_conend() string
}

type IconFactory interface {
	CreateIconFamily() IconFamily
}

// container and leaf should implement this interface
type drawJSON interface {
	Draw()
}

type StyleFamily interface {
	Get_symbol_last() string      // 最后一个子对象的前缀
	Get_symbol() string           // 正常子对象的前缀
	Get_symbol_left_last() string // 左下角

	// prefix
	// 实际上指selfjson的前缀
	Get_prefix_end_g() string   // 全局的最后一个节点，一定在最底部
	Get_prefix_endleaf() string //前json对象的最后一个节点，但下面还有其他的
	Get_prefix() string

	// end
	Get_end() string       // 中间填充
	Get_end_first() string // 右上角
	Get_end_mid() string   // 右边框
	Get_end_last() string  // 右下角
}

type StyleFactory interface {
	CreateStlyeFamily() StyleFamily
}

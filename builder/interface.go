package builder

// 使用生成器模式，可以自定义配置
type drawJSON interface {
	// ParseJSON
	// 根据文件名解析 json
	ParseJSON(jsonfile string) error

	// InitStyle
	// 选择对应的风格
	InitStyle(style string)

	// InitIcon
	// 选择对应的图标族
	InitIcon(IconFamily string)

	// Draw
	// 最后输出
	Draw()
}

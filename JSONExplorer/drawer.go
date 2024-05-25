package JSONExplorer

import (
	"fmt"
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"io"
	"os"
	"strings"
)

type Drawer struct {
	jsonV *jsonvalue.V
	icon  IconFamily
}

// ParseJSON
// 为 Explorer 实现 drawJSON 接口
// 解析输入的 JSON 文件
func (d Drawer) ParseJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	jsonData, err := io.ReadAll(file)

	jsonV, err := jsonvalue.Unmarshal(jsonData)
	d.jsonV = jsonV

	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (d Drawer) InitIcon(icon string) {
	var factory IconFactory
	switch strings.ToLower(icon) {
	case "poker":
		factory = PokerIconFactory{}
	default:
		return
	}
	d.icon = factory.CreateIconFamily()

	fmt.Printf("Leaf icon:%v\nNode icon:%v\n", d.icon.GetLeafIcon(), d.icon.GetNodeIcon())
}

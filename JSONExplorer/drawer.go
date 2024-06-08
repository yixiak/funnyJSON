package JSONExplorer

import (
	"io"
	"os"
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)

type Drawer struct {
	leafIcon  LeafIcon
	nodeIcon  NodeIcon
	Style     StyleFamily
	innerJSON *jsonvalue.V
	iterator  Iterator
	stragegy  Stragegy
}

// ParseJSON
// 为 Explorer 实现 drawJSON 接口
// 解析输入的 JSON 文件
func (d *Drawer) ParseJSON(filename string) error {
	// 读取 json 文件
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	jsonData, _ := io.ReadAll(file)

	jsonV, err := jsonvalue.Unmarshal(jsonData)

	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	d.innerJSON = jsonV

	iter := CreateIterator(d.innerJSON)
	d.iterator = iter
	return nil
}

func (d *Drawer) InitIcon(icon string) {
	var leaffactory LeafIconFactory
	var nodefactory NodeIconFactory
	switch strings.ToLower(icon) {
	case "poker":
		leaffactory = PokerLeafIconFactory{}
		nodefactory = PokerNodeIconFactory{}
	case "defualt":
		leaffactory = DefualtLeafIconFactory{}
		nodefactory = DefualtNodeIconFactory{}
	default:
		return
	}
	d.leafIcon = leaffactory.CreateLeafIcon()
	d.nodeIcon = nodefactory.CreateNodeIcon()
}

func (d *Drawer) InitStyle(style string) {
	var factory StyleFactory
	switch strings.ToLower(style) {
	case "tree":
		factory = &TreeStyleFactory{}
		d.stragegy = &TreeStrategy{}

	case "rec":
		factory = &RecStyleFactory{}
		d.stragegy = &RecStrategy{}
	default:
		return
	}
	d.Style = factory.CreateStlyeFamily()
}

func (d *Drawer) Show() {

	if d.Style == nil {
		d.InitStyle("tree")
		d.stragegy = &TreeStrategy{}
	}

	if d.leafIcon == nil || d.nodeIcon == nil {
		d.InitIcon("defualt")
	}

	d.stragegy.Draw(d.iterator, d.Style, d.leafIcon, d.nodeIcon)

}

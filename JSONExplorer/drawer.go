package JSONExplorer

import (
	"io"
	"log"
	"os"
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)

type Drawer struct {
	icon IconFamily
	root *container
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
	jsonData, err := io.ReadAll(file)

	jsonV, err := jsonvalue.Unmarshal(jsonData)

	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	d.root = &container{
		level:      0,
		innerValue: jsonV,
	}
	return nil
}

func (d *Drawer) InitIcon(icon string) {
	var factory IconFactory
	switch strings.ToLower(icon) {
	case "poker":
		factory = PokerIconFactory{}
	default:
		return
	}
	d.icon = factory.CreateIconFamily()
}

func (d *Drawer) InitStyle(style string) {

}

func (d *Drawer) Show() {

	// 将json对象转为container和leaf对象
	rootChild := []drawJSON{}
	d.root.innerValue.RangeObjectsBySetSequence(func(key string, V *jsonvalue.V) bool {
		var node drawJSON
		switch V.ValueType() {
		case jsonvalue.String:
			node = &leaf{
				key:   key,
				value: V.String(),
			}

		case jsonvalue.Object:
			node = &container{
				key:        key,
				level:      1,
				innerValue: V,
			}

		default:
			log.Fatal("Error when traversing json ")
		}
		rootChild = append(rootChild, node)
		return true
	})

}

type container struct {
	key        string
	level      int
	innerValue *jsonvalue.V
}

type leaf struct {
	key   string
	value string
}

func (c *container) Draw() {}
func (l *leaf) Draw()      {}

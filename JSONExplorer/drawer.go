package JSONExplorer

import (
	"io"
	"os"
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)

type Drawer struct {
	icon  IconFamily
	style StyleFamily
	root  *container
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
	var factory StyleFactory
	switch strings.ToLower(style) {
	case "tree":
		factory = &TreeStyleFactory{}
	case "rec":
		factory = &RecStyleFactory{}
	default:
		return
	}
	d.style = factory.CreateStlyeFamily()
}

func (d *Drawer) Show() {

	if d.style == nil {
		d.InitStyle("tree")
	}

	if d.icon == nil {
		d.InitIcon("poker")
	}

	// 将json对象转为container和leaf对象
	Maxlen := getMaxlen(d.root.innerValue, 0)
	//fmt.Print("Maxlen: ", Maxlen)
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

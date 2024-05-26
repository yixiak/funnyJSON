package JSONExplorer

import (
	"fmt"
	"io"
	"os"
	"strings"

	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
)

type Drawer struct {
	Icon      IconFamily
	Style     StyleFamily
	root      *container
	innerJSON *jsonvalue.V
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
	d.innerJSON = jsonV
	return nil
}

func (d *Drawer) InitIcon(icon string) {
	var factory IconFactory
	switch strings.ToLower(icon) {
	case "poker":
		factory = PokerIconFactory{}
	case "defualt":
		factory = DefualtIconFactory{}
	default:
		return
	}
	d.Icon = factory.CreateIconFamily()
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
	d.Style = factory.CreateStlyeFamily()
}

func (d *Drawer) Show() {

	if d.Style == nil {
		d.InitStyle("tree")
	}

	if d.Icon == nil {
		d.InitIcon("defualt")
	}

	// 将json对象转为container和leaf对象
	Maxlen := getMaxlen(d.innerJSON, 0)
	if Maxlen+2 < 42 {
		Maxlen = 42
	} else {
		Maxlen = Maxlen + 2
	}
	index := 0
	child_len := d.innerJSON.Len()
	for k, v := range d.innerJSON.ForRangeObj() {
		index++
		var prefix, selfjson, end []rune

		if index == 1 {
			prefix = []rune(d.Style.Get_prefix_first())
		} else if index == child_len {
			if v.Len() == 0 {
				prefix = []rune(d.Style.Get_prefix_end_g())
			} else {
				prefix = []rune(d.Style.Get_prefix_endleaf())
			}
		} else {
			prefix_ := []rune(d.Style.Get_prefix())
			prefix = append(prefix, prefix_...)
		}

		if v.ValueType() == jsonvalue.String {
			selfjson = []rune(fmt.Sprintf("%s%s:%s", d.Icon.GetLeaf_Icon(), k, v.String()))
		} else {
			selfjson = []rune(fmt.Sprintf("%s%s", d.Icon.GetNode_Icon(), k))
		}

		end_len := Maxlen - len(prefix) - len(selfjson)
		for i := 0; i < end_len-1; i++ {
			end = append(end, []rune(d.Style.Get_end())...)
		}
		if index == child_len && v.Len() == 0 {
			end = append(end, []rune(d.Style.Get_end_last())...)
		} else if index == 1 {
			end = append(end, []rune(d.Style.Get_end_first())...)
		} else {
			end = append(end, []rune(d.Style.Get_end_mid())...)
		}
		fmt.Printf("%s%s%s\n", string(prefix), string(selfjson), string(end))

		if v.ValueType() == jsonvalue.Object {
			var new_symbol []rune
			if index == child_len {
				new_symbol = append(new_symbol, []rune(d.Style.Get_symbol_last())...)
			} else {
				new_symbol = []rune(d.Style.Get_symbol())
			}
			Draw(d, v, string(new_symbol), Maxlen, index == child_len)
		}
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

func Draw(drawer *Drawer, this *jsonvalue.V, symbol string, maxlen int, is_last bool) {
	child_len := this.Len()
	index := 0
	for k, v := range this.ForRangeObj() {
		var prefix, selfjson, end, my_symbol []rune
		index++
		my_symbol = append(my_symbol, []rune(symbol)...)
		if index == child_len {
			if is_last {
				prefix_end_g := []rune(drawer.Style.Get_prefix_end_g())
				prefix = append(prefix, prefix_end_g...)
				// 替换左下角
				if v.Len() == 0 {
					symbol_left_last := []rune(drawer.Style.Get_symbol_left_last())
					my_symbol = my_symbol[len(symbol_left_last):]
					my_symbol = append(symbol_left_last, my_symbol...)
				}
			} else {
				prefix_end := []rune(drawer.Style.Get_prefix_endleaf())
				prefix = append(prefix, prefix_end...)
			}
		} else {
			prefix_ := []rune(drawer.Style.Get_prefix())
			prefix = append(prefix, prefix_...)
		}

		// 代表是叶子节点
		if v.ValueType() == jsonvalue.String {
			selfjson = []rune(fmt.Sprintf("%s%s:%s", drawer.Icon.GetLeaf_Icon(), k, v))
		} else {
			selfjson = []rune(fmt.Sprintf("%s%s", drawer.Icon.GetNode_Icon(), k))
		}

		end_len := maxlen - len(prefix) - len(selfjson) - len(my_symbol)
		for i := 0; i < end_len-1; i++ {
			end = append(end, []rune(drawer.Style.Get_end())...)
		}
		if is_last && index == child_len && v.Len() == 0 {
			end = append(end, []rune(drawer.Style.Get_end_last())...)
		} else {
			end = append(end, []rune(drawer.Style.Get_end_mid())...)
		}
		fmt.Printf("%s%s%s%s\n", string(my_symbol), string(prefix), string(selfjson), string(end))

		if v.ValueType() == jsonvalue.Object {
			new_symbol := []rune(my_symbol)
			if index == child_len {
				new_symbol = append(new_symbol, []rune(drawer.Style.Get_symbol_last())...)
			} else {
				new_symbol = append(new_symbol, []rune(drawer.Style.Get_symbol())...)
			}
			Draw(drawer, v, string(new_symbol), maxlen, is_last)
		}
	}

}

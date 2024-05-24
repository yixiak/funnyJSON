package builder

import (
	jsonvalue "github.com/Andrew-M-C/go.jsonvalue"
	"io"
	"os"
)

type Drawer struct {
	jsonV *jsonvalue.V
}

// ParseJSON
// 为 builder 实现 drawJSON 接口
// 解析输入的 JSON 文件
func (d *Drawer) ParseJSON(filename string) error {
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

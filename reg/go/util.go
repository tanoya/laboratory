package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

// WriteAppend 将文件追加文件尾巴
func WriteAppend(app string, context []byte) error {
	path := `./report/` + app + ".txt"
	err := ioutil.WriteFile(path, context, os.ModeAppend)
	return err
}

// Format 格式化
func Format(app, method string, costTime int64) string {
	return fmt.Sprintf(" %s %s cost %d nanosecond. \n", app, method, costTime)
}

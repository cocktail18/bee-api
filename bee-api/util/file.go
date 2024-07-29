package util

import (
	"os"
)

func FileExists(filename string) (bool, error) {
	// 使用 os.Stat 函数获取文件信息
	_, err := os.Stat(filename)

	// 判断文件是否存在
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

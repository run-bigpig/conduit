package utils

import (
	"os"
	"path/filepath"
)

// 创建指定路径的文件,如果路径中文件夹不存在则创建
func CreateFile(filePath string) (*os.File, error) {
	//判断文件是否存在
	_, err := os.Stat(filePath)
	if err == nil {
		//文件存在
		return os.Open(filePath)
	}
	//判断是否是文件不存在
	if os.IsNotExist(err) {
		//创建文件夹
		err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return nil, err
		}
		//创建文件
		file, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	//其他错误
	return nil, err
}

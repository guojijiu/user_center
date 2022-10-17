package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// 文件是否存在
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateMutiDir 调用os.MkdirAll递归创建文件夹
func CreateMutiDir(filePath string) error {
	_, err := os.Stat(filePath) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败,error info:", err)
			return err
		}
	}
	return nil
}

func ExportToCSV(filename string, fileData map[int][]string) error {
	if PathExist(filepath.Dir(filename)) == false {
		if err := CreateMutiDir(filepath.Dir(filename)); err != nil {
			return err
		}
	}
	// 不存在则创建;存在则清空;读写模式;
	file, err := os.Create(filename)
	if err != nil {
		_ = fmt.Sprintf("打开文件失败：%s，失败原因：%s", filename, err)
		return err
	}
	// 延迟关闭
	defer file.Close()

	// 写入UTF-8 BOM，防止中文乱码
	//file.WriteString("xEFxBBxBF")

	w := csv.NewWriter(file)

	// 按照key排序
	var keys []int
	for k := range fileData {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		w.Write(fileData[key])
		// 刷新缓冲
		w.Flush()
	}

	return nil
}

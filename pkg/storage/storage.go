package storage

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"user_center/config"
	"user_center/pkg/randc"
)

var Storage storage

type storage struct {
	// public 目录的文件夹路径
	AbsPath  string
	DiskName string
}

func Init(storagePath string) {
	Storage = storage{
		AbsPath:  filepath.Join(storagePath, config.Filesystems.Disks.Local.Root),
		DiskName: config.Filesystems.Default,
	}
}

func (s *storage) Disk(diskName string) *storage {
	s.DiskName = diskName
	return s
}

func (s *storage) FullPath(path string) string {
	return filepath.Join(s.AbsPath, path)
}

/**
判断文件是否存在, 使用的是文件的保存路径 storage/app/public/ 下
*/
func (s *storage) Exists(path string) bool {
	fullPath := s.FullPath(path)
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
保存base64字符串为文件
*/
func (s *storage) StoreBase64RandomName(path string, ImgBase64 string) (string, error) {
	b64data := ImgBase64[strings.IndexByte(ImgBase64, ',')+1:]
	fileContent, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return "", err
	}
	fullPath := s.FullPath(path)
	fmt.Printf("full: %s \n", fullPath)
	fileName := fmt.Sprintf("%s.jpg", randc.UUID())
	filePath := filepath.Join(fullPath, fileName)
	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(filePath, fileContent, 0777)
	if err != nil {
		return "", err
	}

	return filepath.Join(path, fileName), nil
}

/**
文件转base64
*/
func (s *storage) FileToBase64(path string) (string, error) {
	fileContent, err := ioutil.ReadFile(s.FullPath(path))
	if err != nil {
		return "", err
	}
	base64String := base64.StdEncoding.EncodeToString(fileContent)
	return base64String, nil
}

package storage_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
	"user_center/app"
	boot "user_center/boot"
	"user_center/pkg/storage"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestStorageSaveBase64(t *testing.T) {
	frontImgBase64, _ := ioutil.ReadFile("./testdata/id-card-back-base64.txt")
	savePath, err := storage.Storage.StoreBase64RandomName(fmt.Sprintf("id_card_imgs/%s", "123456"), string(frontImgBase64))
	log.Println(savePath, err)
	assert.Nil(t, err)
	assert.NotEqual(t, "", savePath)
	assert.True(t, storage.Storage.Exists(savePath))
}

func TestGetStorageAbsPath(t *testing.T) {
	log.Println(storage.Storage.AbsPath)
}

func TestFileToBase64(t *testing.T) {
	frontImgBase64, _ := ioutil.ReadFile(filepath.Join(app.TestPath, "testdata", "id-card-back-base64.txt"))
	savePath, err := storage.Storage.StoreBase64RandomName(fmt.Sprintf("id_card_imgs/%s", "123456"), string(frontImgBase64))
	log.Println(savePath, err)
	assert.Nil(t, err)
	assert.NotEqual(t, "", savePath)
	assert.True(t, storage.Storage.Exists(savePath))

	base64String, err := storage.Storage.FileToBase64(savePath)
	assert.Nil(t, err)
	assert.NotEqual(t, "", base64String)
	log.Println(base64String)
}

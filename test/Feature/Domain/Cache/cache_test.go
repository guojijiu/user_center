package user_test

import (
	"fmt"
	"testing"
	"user_center/app/Domain/Cache"
	"user_center/boot"
	"user_center/pkg/tool"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

// go test -v test/Feature/Domain/Cache/cache_test.go -test.run TestStore
func TestStore(t *testing.T) {
	codeCache := &Cache.CodeCache{}
	str := tool.RandomNumber(6)
	res, err := codeCache.SetCacheKey("user:register_code", "123@qq.com").Store(str)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("success:", res)
}

// go test -v test/Feature/Domain/Cache/cache_test.go -test.run TestGet
func TestGet(t *testing.T) {
	codeCache := &Cache.CodeCache{}
	res, err := codeCache.SetCacheKey("user:register_code", "123@qq.com").Get()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("success:", res)
}

// go test -v test/Feature/Domain/Cache/cache_test.go -test.run TestDelete
func TestDelete(t *testing.T) {
	codeCache := &Cache.CodeCache{}
	res, err := codeCache.SetCacheKey("user:register_code", "123@qq.com").Delete()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("success:", res)
}

package Organize_test

import (
	"fmt"
	genid "github.com/srlemon/gen-id"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/DetailClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ForbiddenClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ListClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/StoreClient"
	"user_center/app/Http/Controllers/API/Admin/Context/Client/UpdateClient"
	"user_center/boot"
	"user_center/pkg/test"
)

var httptest *test.Http

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	httptest = test.New(app.GetEngineRouter())
	m.Run()
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestStore
func TestStore(t *testing.T) {
	w := httptest.Post("/api/admin/organize/store", StoreClient.Req{
		Name:   genid.NewGeneratorData().Name,
		Mark:   genid.NewGeneratorData().Name,
		Remark: genid.NewGeneratorData().Address,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestDetail
func TestDetail(t *testing.T) {
	resp := httptest.Get("/api/admin/organize/detail", DetailClient.Req{
		ID: 1,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestList
func TestList(t *testing.T) {
	resp := httptest.Get("/api/admin/organize/list", ListClient.Req{
		Page: 1,
		Size: 2,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestUpdate
func TestUpdate(t *testing.T) {
	w := httptest.Call("PUT", "/api/admin/organize/update", UpdateClient.Req{
		ID:     1,
		Name:   genid.NewGeneratorData().Name,
		Mark:   genid.NewGeneratorData().Name,
		Remark: genid.NewGeneratorData().Address,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestForbidden
func TestForbidden(t *testing.T) {
	w := httptest.Call("POST", "/api/admin/organize/forbidden", ForbiddenClient.Req{
		ID:          1,
		IsForbidden: 1,
	})
	fmt.Println(w.Body)
}

package Organize_test

import (
	"fmt"
	genid "github.com/srlemon/gen-id"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/DetailOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ForbiddenOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/ListOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/StoreOrganize"
	"user_center/app/Http/Controllers/API/Admin/Context/Organize/UpdateOrganize"
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
	w := httptest.Post("/api/admin/organize/store", StoreOrganize.Req{
		Name:   genid.NewGeneratorData().Name,
		Mark:   genid.NewGeneratorData().Name,
		Remark: genid.NewGeneratorData().Address,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestDetail
func TestDetail(t *testing.T) {
	resp := httptest.Get("/api/admin/organize/detail", DetailOrganize.Req{
		ID: 1,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestList
func TestList(t *testing.T) {
	resp := httptest.Get("/api/admin/organize/list", ListOrganize.Req{
		Page: 1,
		Size: 2,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestUpdate
func TestUpdate(t *testing.T) {
	w := httptest.Call("PUT", "/api/admin/organize/update", UpdateOrganize.Req{
		ID:     1,
		Name:   genid.NewGeneratorData().Name,
		Mark:   genid.NewGeneratorData().Name,
		Remark: genid.NewGeneratorData().Address,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Organize/organize_test.go -test.run TestForbidden
func TestForbidden(t *testing.T) {
	w := httptest.Call("POST", "/api/admin/organize/forbidden", ForbiddenOrganize.Req{
		ID:          1,
		IsForbidden: 1,
	})
	fmt.Println(w.Body)
}

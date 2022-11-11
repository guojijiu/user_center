package Organize_test

import (
	"fmt"
	genid "github.com/srlemon/gen-id"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/DetailDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ForbiddenDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/ListDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/StoreDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/UpdateDepartment"
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

// go test -v test/Feature/Admin/Department/department_test.go -test.run TestStore
func TestStore(t *testing.T) {
	w := httptest.Post("/api/admin/department/store", StoreDepartment.Req{
		Name:   genid.NewGeneratorData().Name,
		Mark:   genid.NewGeneratorData().Name,
		Remark: genid.NewGeneratorData().Address,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Department/department_test.go -test.run TestDetail
func TestDetail(t *testing.T) {
	resp := httptest.Get("/api/admin/department/detail", DetailDepartment.Req{
		ID: 1,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Department/department_test.go -test.run TestList
func TestList(t *testing.T) {
	resp := httptest.Get("/api/admin/department/list", ListDepartment.Req{
		Page: 1,
		Size: 2,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Department/department_test.go -test.run TestUpdate
func TestUpdate(t *testing.T) {
	w := httptest.Call("PUT", "/api/admin/department/update", UpdateDepartment.Req{
		ID:     1,
		Name:   genid.NewGeneratorData().Name,
		Mark:   genid.NewGeneratorData().Name,
		Remark: genid.NewGeneratorData().Address,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Department/department_test.go -test.run TestForbidden
func TestForbidden(t *testing.T) {
	w := httptest.Call("POST", "/api/admin/department/forbidden", ForbiddenDepartment.Req{
		ID:          1,
		IsForbidden: 1,
	})
	fmt.Println(w.Body)
}

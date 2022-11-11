package role_test

import (
	"fmt"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/BindDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/BindPermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DeleteRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DetailRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/GetBindDepartment"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/ListRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/StoreRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/UpdateRole"
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

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestStore
func TestStore(t *testing.T) {
	w := httptest.Post("/api/admin/role/store", StoreRole.Req{
		ClientID: 1,
		Name:     "aaa",
		Sort:     2,
		Mark:     "233",
		Remark:   "445",
	})
	fmt.Println(w.Body)

}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestDetail
func TestDetail(t *testing.T) {

	resp := httptest.Get("/api/admin/role/detail", DetailRole.Req{
		ID: 2,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestList
func TestList(t *testing.T) {
	resp := httptest.Get("/api/admin/role/list", ListRole.Req{
		Page: 1,
		Size: 10,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestUpdate
func TestUpdate(t *testing.T) {
	w := httptest.Call("PUT", "/api/admin/role/update", UpdateRole.Req{
		ID:     1,
		Name:   "33",
		Sort:   44,
		Mark:   "5",
		Remark: "66",
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestDelete
func TestDelete(t *testing.T) {
	w := httptest.Call("DELETE", "/api/admin/role/delete", DeleteRole.Req{
		ID: 1,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestBindPermission
func TestBindPermission(t *testing.T) {

	w := httptest.Post("/api/admin/role/bind", BindPermission.Req{
		ID:            1,
		PermissionIDs: []uint{6, 7, 8, 9, 10},
	})
	fmt.Println(w.Body)

}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestBindDepartment
func TestBindDepartment(t *testing.T) {
	w := httptest.Call("POST", "/api/admin/role/bind_department", BindDepartment.Req{
		ID:            1,
		DepartmentIDs: []uint{1, 2},
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Role/role_test.go -test.run TestGetBindDepartment
func TestGetBindDepartment(t *testing.T) {

	resp := httptest.Get("/api/admin/role/bind_department", GetBindDepartment.Req{
		ID: 1,
	})
	fmt.Println(resp.Body)
}

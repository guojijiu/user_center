package role_test

import (
	"fmt"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DeleteRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DetailRole"
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

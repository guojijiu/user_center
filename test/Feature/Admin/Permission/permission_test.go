package permission_test

import (
	"fmt"
	"github.com/bxcodec/faker/v4"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/DeletePermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Permission/GetTreePermission"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/DetailRole"
	"user_center/app/Http/Controllers/API/Admin/Context/Role/ListRole"
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

// go test -v test/Feature/Admin/Permission/permission_test.go -test.run TestStore
func TestStore(t *testing.T) {
	type Store struct {
		ClientID   uint   `faker:"boundary_start=1, boundary_end=3" json:"client_id"`
		Name       string `faker:"lang=chi,len=5" json:"name"`
		Sort       uint   `faker:"boundary_start=1, boundary_end=999" json:"sort"`
		Mark       string `faker:"word" json:"mark"`
		Type       uint   `faker:"oneof: 1,2,3,4" json:"type"`
		ParentID   uint   `faker:"boundary_start=1, boundary_end=9999" json:"parent_id"`
		Remark     string `faker:"word" json:"remark"`
		IconPath   string `faker:"url" json:"icon_path"`
		RouteName  string `faker:"word" json:"route_name"`
		RoutePath  string `faker:"url" json:"route_path"`
		ModulePath string `faker:"url" json:"module_path"`
		HiddenAt   string `faker:"timestamp" json:"hidden_at"`
	}
	var store Store
	_ = faker.FakeData(&store)

	w := httptest.Post("/api/admin/permission/store", store)
	fmt.Println(w.Body)

}

// go test -v test/Feature/Admin/Permission/permission_test.go -test.run TestDetail
func TestDetail(t *testing.T) {

	resp := httptest.Get("/api/admin/permission/detail", DetailRole.Req{
		ID: 2,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Permission/permission_test.go -test.run TestList
func TestList(t *testing.T) {
	resp := httptest.Get("/api/admin/permission/list", ListRole.Req{
		Page: 1,
		Size: 10,
	})
	fmt.Println(resp.Body)
}

// go test -v test/Feature/Admin/Permission/permission_test.go -test.run TestUpdate
func TestUpdate(t *testing.T) {
	type Update struct {
		ID            uint   `faker:"oneof: 1" json:"id"`
		Name          string `faker:"lang=chi,len=5" json:"name"`
		Sort          uint   `faker:"boundary_start=1, boundary_end=999" json:"sort"`
		Mark          string `faker:"word" json:"mark"`
		Type          uint   `faker:"oneof: 1,2,3,4" json:"type"`
		ParentID      uint   `faker:"boundary_start=1, boundary_end=9999" json:"parent_id"`
		Remark        string `faker:"word" json:"remark"`
		IconPath      string `faker:"url" json:"icon_path"`
		RouteName     string `faker:"word" json:"route_name"`
		RoutePath     string `faker:"url" json:"route_path"`
		ModulePath    string `faker:"url" json:"module_path"`
		RequestMethod string `faker:"oneof: POST, GET, DELETE, PUT" json:"request_method"`
		HiddenAt      string `faker:"timestamp" json:"hidden_at"`
	}
	var update Update
	_ = faker.FakeData(&update)
	w := httptest.Call("PUT", "/api/admin/permission/update", update)
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Permission/permission_test.go -test.run TestDelete
func TestDelete(t *testing.T) {
	w := httptest.Call("DELETE", "/api/admin/permission/delete", DeletePermission.Req{
		ID: 1,
	})
	fmt.Println(w.Body)
}

// go test -v test/Feature/Admin/Permission/permission_test.go -test.run TestGetTree
func TestGetTree(t *testing.T) {
	w := httptest.Get("/api/admin/permission/tree", GetTreePermission.Req{
		ClientID: 1,
	})
	fmt.Println(w.Body)
}

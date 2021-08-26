package user_test

import (
	"encoding/json"
	"fmt"
	"github.com/srlemon/gen-id"
	"github.com/stretchr/testify/assert"
	"testing"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/User/DetailUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ListUser"
	"user_center/app/Http/Controllers/API/Admin/Context/User/StoreUser"
	"user_center/app/Http/Controllers/API/Admin/Responses"
	"user_center/app/Model"
	"user_center/boot"
	"user_center/pkg/db"
	"user_center/pkg/test"
)

var httptest *test.Http

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	httptest = test.New(app.GetEngineRouter())
	m.Run()
}

// go test -v test/Feature/Admin/User/user_test.go -test.run TestRegister
func TestRegister(t *testing.T) {
	w := httptest.Post("/api/admin/user/store", StoreUser.Req{
		Account:  genid.NewGeneratorData().Name,
		Phone:    genid.NewGeneratorData().PhoneNum,
		Email:    genid.NewGeneratorData().Email,
		Passwd:   "123456",
		Nickname: "zhangsan",
		Birthday: "2021-11-12 00:00:00",
	})
	fmt.Println(w.Body)
	//t.Logf("resp: %s", w.Body)
	//assert.Equal(t, w.Code, 200)
	//r := Responses.Response{}
	//err = json.Unmarshal(w.Body.Bytes(), &r)
	//assert.Nil(t, err)
	//assert.Equal(t, 0, r.Code)

}

// go test -v test/Feature/Admin/User/user_test.go -test.run TestDetail
func TestDetail(t *testing.T) {

	resp := httptest.Get("/api/admin/user/detail", DetailUser.Req{
		ID: 8,
	})
	fmt.Println(resp.Body)
}

func TestFindPasswordToken(t *testing.T) {
	user := &Model.UserAuth{}
	err := db.Def().First(&user).Error
	assert.Nil(t, err)
	assert.NotEmpty(t, user.Phone)
	resp := httptest.Get("/api/auth/find/password/token", StoreUser.Req{
		Phone: user.Phone,
	})
	t.Logf("resp: %s", resp.Body)
	assert.Equal(t, resp.Code, 200)
	r := Responses.Response{}
	err = json.Unmarshal(resp.Body.Bytes(), &r)
	if body, ok := r.Body.(map[string]interface{}); !ok {
		t.Error("响应处理失败", body)
		t.FailNow()
	} else {
		assert.NotEmpty(t, body["find_password_token"])
	}
}

// go test -v test/Feature/Admin/User/user_test.go -test.run TestList
func TestList(t *testing.T) {
	resp := httptest.Get("/api/admin/user/list", ListUser.Req{
		Page: 1,
		Size: 2,
	})
	fmt.Println(resp.Body)
}

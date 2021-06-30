package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMobileValidator(t *testing.T) {
	var err error
	type user struct {
		Mobile string `json:"mobile" validate:"required,mobile" comment:"手机号"`
	}
	u := user{
		Mobile: "",
	}

	err = Validate.Struct(u)
	//assert.NotNil(t, err)
	if err != nil {
		t.Logf("err: %s", err.Error())
	}
	u.Mobile = "13517210601"
	err = Validate.Struct(u)
	assert.Nil(t, err)
}

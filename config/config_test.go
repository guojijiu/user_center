package config_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"user_center/boot"
	"user_center/config"
)

func TestVal(t *testing.T) {
	boot.SetInTest()
	assert.NotEmpty(t, config.Name)
	assert.NotEmpty(t, os.Getenv("APP_NAME"))
	assert.Equal(t, config.Name, os.Getenv("APP_NAME"))
	assert.Equal(t, config.APPKey, os.Getenv("APP_KEY"))
}

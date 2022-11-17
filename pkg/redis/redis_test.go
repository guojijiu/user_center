package redis_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"user_center/pkg/redis"
)

// go test -v pkg/redis/redis_test.go -test.run TestConn
func TestConn(t *testing.T) {
	var err error
	_, err = redis.InitDef()
	assert.Nil(t, err)
	err = redis.Def().Ping().Err()
	assert.Nil(t, err)
}

// go test -v pkg/redis/redis_test.go -test.run TestQuery
func TestQuery(t *testing.T) {
	var err error
	_, err = redis.InitDef()
	assert.Nil(t, err)
	//assert.Nil(t, db.Def().DB().Ping())
}

// go test -v pkg/redis/redis_test.go -test.run TestSet
func TestSetStr(t *testing.T) {
	var err error
	err = redis.Def().Set("cccc", "1111", 100*time.Second).Err()
	assert.Nil(t, err)
}

// go test -v pkg/redis/redis_test.go -test.run TestGetStr
func TestGetStr(t *testing.T) {
	a := redis.Def().Get("dddd")
	fmt.Println(a)
}

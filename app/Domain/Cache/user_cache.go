package Cache

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"user_center/config"
)

type UserCache struct {
	userKey string
}

type UserInfo struct {
	ID      uint   `comment:"主键ID" json:"id"`
	Account string `comment:"账号" json:"account"`
	Phone   string `comment:"手机号" json:"phone"`
	Email   string `comment:"邮箱" json:"email"`
}

func (uc *UserCache) SetCacheKey(userID uint) *UserCache {
	uc.userKey = fmt.Sprintf("%suser:info:%s", config.CachePrefix, strconv.Itoa(int(userID)))
	return uc
}

func (uc *UserCache) Store(userInfo *UserInfo) (bool, error) {
	data, _ := json.Marshal(&userInfo)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	Cache.HMSet(uc.userKey, m)
	return Cache.Expire(uc.userKey, 6*time.Hour).Result()
}

func (uc *UserCache) SetFieldByID() {

}

func (uc *UserCache) Get() UserInfo {
	var userInfo UserInfo
	data, _ := Cache.HGetAll(uc.userKey).Result()
	arr, _ := json.Marshal(data)
	_ = json.Unmarshal(arr, &userInfo)
	return userInfo
}

func (uc *UserCache) Delete() (int64, error) {
	return Cache.Del(uc.userKey).Result()
}

func (uc *UserCache) GetTTL() {

}

func (uc *UserCache) RefreshExpire() {

}

func (uc *UserCache) Exist() (int64, error) {
	return Cache.Exists(uc.userKey).Result()
}

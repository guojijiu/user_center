package Cache

import (
	"fmt"
	"user_center/config"
)

type UserCache struct {
	userKey string
}

func (uc *UserCache) SetCacheKey(userID uint) *UserCache {
	uc.userKey = fmt.Sprintf("%suser:info:%s", config.CachePrefix, userID)
	return uc
}

func (uc *UserCache) Store() {

}

func (uc *UserCache) SetFieldByID() {

}

func (uc *UserCache) Get() {

}

func (uc *UserCache) Delete() {

}

func (uc *UserCache) GetTTL() {

}

func (uc *UserCache) RefreshExpire() {

}

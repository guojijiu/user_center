package Cache

import (
	"fmt"
	"time"
	"user_center/config"
)

type CodeCache struct {
	codeKey string
}

func (cc *CodeCache) SetCacheKey(strPrefix string, key string) *CodeCache {
	cc.codeKey = fmt.Sprintf("%s%s:%s", config.CachePrefix, strPrefix, key)
	return cc
}

func (cc *CodeCache) Store(numStr string) (string, error) {
	return Cache.Set(cc.codeKey, numStr, 30*time.Minute).Result()
}

func (cc *CodeCache) Get() (string, error) {
	return Cache.Get(cc.codeKey).Result()
}

func (cc *CodeCache) Delete() (int64, error) {
	return Cache.Del(cc.codeKey).Result()
}

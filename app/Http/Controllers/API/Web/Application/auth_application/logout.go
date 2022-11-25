package auth_application

import (
	"user_center/app/Domain/Cache"
)

func LogoutUser(id uint) error {
	cache := Cache.UserCache{}
	_, _ = cache.SetCacheKey(id).Delete()
	return nil
}

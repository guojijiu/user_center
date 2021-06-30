package encryption

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 用bcrypt算法将明文hash
func BcryptHash(raw string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	return string(hashBytes), err
}

// BcryptCheck 用bcrypt算法比对密码明文和hash值是否匹配
func BcryptCheck(raw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(raw))
	return err == nil
}

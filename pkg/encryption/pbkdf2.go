package encryption

import (
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/pbkdf2"
	"hash"
	"strings"
)

type pbkdf2options struct {
	passwdLen       int
	iterationsCount int
	pSalt           *saltType
	hashfunc        func() hash.Hash
}

type saltType struct {
	saltLen   int
	saltValue []byte
}

func NewSalt(saltLen int, allowChars ...[]byte) *saltType {
	return &saltType{
		saltLen:   saltLen,
		saltValue: RandBytes(saltLen, allowChars...),
	}
}

var DefaultPBKDF2Options = pbkdf2options{
	passwdLen:       24,
	iterationsCount: 10000,
	hashfunc:        sha1.New,
}

// GeneratePasswdPBKDF2Key 用pbkdf2生成密码
func (pbkdf2opt *pbkdf2options) GeneratePasswdPBKDF2Key(rawtext, saltbytes []byte) (string, string) {
	if len(saltbytes) == 0 {
		saltbytes = NewSalt(10, []byte("ABCDEFGHJKLMNPQRSTUVWXY23456789ABCDEFGHJKLMNPQRSTUVWXY23456789abcdefghjkmnpqrstuvwxy23456789abcdefghjkmnpqrstuvwxy23456789")).saltValue
	}

	dk := pbkdf2.Key(rawtext, saltbytes, pbkdf2opt.iterationsCount, pbkdf2opt.passwdLen, pbkdf2opt.hashfunc)

	return hex.EncodeToString(dk), string(saltbytes)
}

// CheckPBKDF2Passwd 校验通过pbkdf2生成的密码是否正确
func (pbkdf2options *pbkdf2options) CheckPBKDF2Passwd(rawtext, passwd, salt string) bool {
	passwd2, _ := pbkdf2options.GeneratePasswdPBKDF2Key([]byte(rawtext), []byte(salt))
	return strings.EqualFold(passwd2, passwd)
}

// CheckPBKDF2PasswdForVzhuoTaskSYS 专用于微桌任务系统的登录密码校验
func (pbkdf2options *pbkdf2options) CheckPBKDF2PasswdForVzhuoTaskSYS(rawtext, passwd, salt string) bool {

	if len(salt) > 15 {
		return MD5(salt+rawtext) == passwd
	}

	passwd2, _ := pbkdf2options.GeneratePasswdPBKDF2Key([]byte(rawtext), []byte(salt))
	return strings.EqualFold(passwd2, passwd)
}

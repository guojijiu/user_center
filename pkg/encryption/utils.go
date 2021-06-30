package encryption

import (
	"bytes"
	"crypto/md5"
	crand "crypto/rand"
	"encoding/hex"
	"math/big"
)

// PKCS7Padding 将rawtext按照blocksize的整数倍补码
func PKCS7Padding(rawtext []byte, blockSize int) []byte {
	padding := blockSize - len(rawtext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(rawtext, padText...)
}

// PKCS7Padding 去除rawtext的补码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	//fmt.Println(length)
	unpadding := int(origData[length-1])
	//fmt.Println(origData[length-1])
	//fmt.Println(unpadding)
	return origData[:(length - unpadding)]
}

func RandBytes(n int, allowedChars ...[]byte) []byte {
	var defaultLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var letters []byte

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]byte, n)
	for i := range b {
		theN, _ := crand.Int(crand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[theN.Int64()]
	}
	return b
}

func RandStr(n int, allowedChars ...[]byte) string {
	return string(RandBytes(n, allowedChars...))
}

// MD5
func MD5(data string) string {
	md5Obj := md5.New()
	md5Obj.Write([]byte(data))
	return hex.EncodeToString(md5Obj.Sum(nil))
}

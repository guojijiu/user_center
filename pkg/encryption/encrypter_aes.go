package encryption

import (
	"crypto/aes"
	crypt_cipher "crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"user_center/config"
)

const (
	KEY_STR_PREFIX     = "base64:"
	CIPHER_AES_128_CBC = "AES-128-CBC"
	CIPHER_AES_256_CBC = "AES-256-CBC"

	KEY_LEN_FOR_AES_128 = 16
	KEY_LEN_FOR_AES_256 = 32
)

var defaultCipherType string = CIPHER_AES_256_CBC

func SetDefaultCipherType(cipherType string) {
	defaultCipherType = cipherType
}

type encryptor struct {
	// key
	keyStr string

	currKeyLen int
	maxkeyLen  int

	keyBytes []byte

	// cipher: the algorithm used for encryption
	cipher string

	cipherAgent crypt_cipher.Block

	keyBlockSize int
}

var DefaultEncryptor, _ = NewEncryptor(config.APPKey, CIPHER_AES_256_CBC)

func NewEncryptor(keyStr, cipher string) (*encryptor, error) {
	keyBytes, err := ParseKeyStrToBinaryByte(keyStr)
	//fmt.Printf("keyBytes=%s\n", string(keyBytes))
	if err != nil {
		return nil, err
	}

	pEncryptor := &encryptor{
		keyStr:     keyStr,
		currKeyLen: len(keyBytes),
		keyBytes:   keyBytes,
	}

	switch pEncryptor.currKeyLen {
	default:
		return nil, aes.KeySizeError(pEncryptor.currKeyLen)
	case 16, 24, 32:
		break
	}
	// 实例化一个加解密处理器。并做了密钥分组
	pEncryptor.cipherAgent, err = aes.NewCipher(pEncryptor.keyBytes)
	if err != nil {
		return nil, err
	}
	pEncryptor.keyBlockSize = pEncryptor.cipherAgent.BlockSize()

	switch cipher {
	default:
		fallthrough
	case CIPHER_AES_256_CBC:
		pEncryptor.maxkeyLen = KEY_LEN_FOR_AES_256
	case CIPHER_AES_128_CBC:
		pEncryptor.maxkeyLen = KEY_LEN_FOR_AES_128
	}

	return pEncryptor, nil
}

// EncryptWithCBC 以AES-CBC算法加密
func (encr *encryptor) EncryptWithCBC(raw string) string {
	rawBytes := []byte(raw)

	//fmt.Printf("待加密字节数组：%v\n", rawBytes)

	rawBytes = PKCS7Padding(rawBytes, encr.keyBlockSize)

	//fmt.Printf("待加密补码后：%v\n", rawBytes)

	cbcCipher := crypt_cipher.NewCBCEncrypter(encr.cipherAgent, encr.keyBytes[:encr.keyBlockSize])
	crypted := make([]byte, len(rawBytes))
	cbcCipher.CryptBlocks(crypted, rawBytes)

	return base64.StdEncoding.EncodeToString(crypted)
}

// Decrypt 以AES-CBC算法解密
func (encr *encryptor) DecryptWithCBC(crypted string) string {
	cryptedBytes, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		panic(err.Error())
	}
	cbcCipher := crypt_cipher.NewCBCDecrypter(encr.cipherAgent, encr.keyBytes[:encr.keyBlockSize])
	rawBytes := make([]byte, len(cryptedBytes))
	cbcCipher.CryptBlocks(rawBytes, cryptedBytes)

	//fmt.Printf("解密后还未去补码：%v\n", rawBytes)

	rawBytes = PKCS7UnPadding(rawBytes)

	return string(rawBytes)
}

func GetKeyGoodLen(cipherType string) int {
	switch cipherType {
	default:
		fallthrough
	case CIPHER_AES_256_CBC:
		return KEY_LEN_FOR_AES_256
	case CIPHER_AES_128_CBC:
		return KEY_LEN_FOR_AES_128
	}
}

// GenerateKey 生成字节格式的加解密key
func GenerateKey() []byte {
	return GenerateRandBytesWithCrypto(GetKeyGoodLen(defaultCipherType))
}

// GenerateBase64Key 生成一个以base64编码的加解密key并带上指定的前缀
func GenerateBase64Key() (string, error) {
	keyBytes := GenerateKey()
	if len(keyBytes) == 0 {
		return "", errors.New("Generate app random key failed.")
	}
	return KEY_STR_PREFIX + base64.StdEncoding.EncodeToString(keyBytes), nil
}

// ParseKeyStrToBinaryByte 将base64编码后且以指定字符串为前缀的key解析为字节切片
func ParseKeyStrToBinaryByte(keyStr string) (keyBytes []byte, err error) {
	if len(keyStr) == 0 {
		err = errors.New("encryptor key length can not be 0")
		return
	}
	if strings.HasPrefix(keyStr, KEY_STR_PREFIX) {
		keyStrAfter := keyStr[len(KEY_STR_PREFIX):]
		//fmt.Println("keyStrAfter=", keyStrAfter)
		if len(keyStrAfter) == 0 {
			err = errors.New(fmt.Sprintf("encryptor key error: %s must be prefix", KEY_STR_PREFIX))
		}
		//fmt.Println(keyStrAfter)
		keyBytes, err = base64.StdEncoding.DecodeString(keyStrAfter)
		if err != nil {
			return
		}
	}
	if len(keyBytes) == 0 {
		err = errors.New("Parse key string to binary byte failed.")
	}

	return
}

// GenerateRandBytesWithCrypto 生成指定长度指定字符的随机字节切片
func GenerateRandBytesWithCrypto(n int, allowedChars ...[]byte) []byte {
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

package encryption

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"user_center/config"
	"user_center/pkg/tool"
)

var commonInput = []byte{
	0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
	0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
	0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
	0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
}

func TestStringSplitAfter(t *testing.T) {
	testCaseV := "QgvICXZ5WHDgL6GSuBN7RTN/QvVt1Z9lxd0GGPcVvhM="
	testCaseFV := KEY_STR_PREFIX + testCaseV
	s := strings.SplitAfter(testCaseFV, KEY_STR_PREFIX)
	tool.Dump(s)
	assert.Equal(t, testCaseV, s[1])
}

func TestParseKeyStrToBinaryByte(t *testing.T) {
	testCaseV := "QgvICXZ5WHDgL6GSuBN7RTN/QvVt1Z9lxd0GGPcVvhM="
	testCaseFV := KEY_STR_PREFIX + testCaseV
	//fmt.Println(testCaseFV)
	r, err := ParseKeyStrToBinaryByte(testCaseFV)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%v\n", r)
}

func TestEncryptor_EncryptWithCBC(t *testing.T) {
	testCases := []string{
		"hello,world!",
		"王磊",
		"明朗",
		"詹光",
		"刘伟",
		"LoveCoder!结算系统小组LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!",
		"哇娃瓦挖袜哇蛙洼凹🐸",
		string(commonInput),
	}
	keyStr := config.APPKey
	//keyStr := KEY_STR_PREFIX + base64.StdEncoding.EncodeToString([]byte("12345678901234567890123456789012"))
	cipher, err := NewEncryptor(keyStr, CIPHER_AES_256_CBC)
	if err != nil {
		t.Fatal(err)
	}
	for _, caseV := range testCases {
		cryptedStr := cipher.EncryptWithCBC(caseV)
		fmt.Printf("加密后：[%s]\n", cryptedStr)
		if caseV != cipher.DecryptWithCBC(cryptedStr) {
			t.Errorf("[%s] 经过加密后解密不能还原\n", caseV)
		}
	}
}

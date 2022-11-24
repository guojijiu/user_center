package Encryption_test

import (
	"crypto/aes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"user_center/pkg/encryption"
)

var testCases = []struct {
	rawByte    []byte
	padRawByte []byte
}{
	{
		[]byte("aaa"),
		[]byte{119, 97, 110, 103, 108, 101, 105, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	},
	{
		[]byte("哇娃瓦挖袜哇蛙洼凹🐸"),
		[]byte{0xe5, 0x93, 0x87, 0xe5, 0xa8, 0x83, 0xe7, 0x93, 0xa6, 0xe6, 0x8c, 0x96, 0xe8, 0xa2, 0x9c, 0xe5, 0x93, 0x87, 0xe8, 0x9b, 0x99, 0xe6, 0xb4, 0xbc, 0xe5, 0x87, 0xb9, 0xf0, 0x9f, 0x90, 0xb8, 0x1},
	},
}

func TestPKCS7Padding(t *testing.T) {
	for _, testV := range testCases {
		result := encryption.PKCS7Padding(testV.rawByte, aes.BlockSize)
		fmt.Println(result)
		assert.Equal(t, testV.padRawByte, result)
	}
}

func TestPKCS7UnPadding(t *testing.T) {
	testV := []byte{0xe5, 0x93, 0x87, 0xe5, 0xa8, 0x83, 0xe7, 0x93, 0xa6, 0xe6, 0x8c, 0x96, 0xe8, 0xa2, 0x9c, 0xe5, 0x93, 0x87, 0xe8, 0x9b, 0x99, 0xe6, 0xb4, 0xbc, 0xe5, 0x87, 0xb9, 0xf0, 0x9f, 0x90, 0xb8, 0x1}
	result := encryption.PKCS7UnPadding(testV)
	fmt.Println(result)
}

func TestRandStr(t *testing.T) {
	result := encryption.RandStr(10, []byte("ABCDEFGHJKLMNPQRSTUVWXY23456789ABCDEFGHJKLMNPQRSTUVWXY23456789abcdefghjkmnpqrstuvwxy23456789abcdefghjkmnpqrstuvwxy23456789"))
	fmt.Println(result)
}

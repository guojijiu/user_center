package encryption

import (
	"encoding/base64"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenRSAkey(t *testing.T) {
	err := GenRSAkey(2048, false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnDecryptWithRSA(t *testing.T) {
	testCases := [][]byte{
		[]byte("王磊橘子小豆🍊豆豆小豆几"),
		[]byte("hello,world!"),
		[]byte("LoveCoder!结算系统小组LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!"),
	}
	for _, testV := range testCases {
		cipherText, err := EncryptWithRSA(testV)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("[%s] 加密后得到 [%s] \n", testV, base64.StdEncoding.EncodeToString(cipherText))

		raw, err := DecryptWithRSA(cipherText)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, raw, testV)
	}
}

func TestDecryptWithRSA(t *testing.T) {
	testCases := []struct {
		cipherText string
		rawText    string
	}{
		//{"MTdhOTc2NmFjZTdhODkxN2RhNmFiNjRkZjZlZTU5ZTkwYjlmMTJjMzVlZjI5MjkzN2MxMzQ5YjFjNTU0MjczYjBhY2ZhMjVhYzVmZjQ4MTU3YzY1OGQwYzVjZDYyODNiMDFkMjgwZmFiYjdjOWQxZWRhMzBkZmRjYTdiNmYxYjM3ZjQwOTRlNGNiMGI5ZDM5N2YxYzdkNmEwMjBiODAzMmFiOTAwYmZlNzYyNjUwOGM5MTQ5OGViYWJiNTYzMzM3NTEyNzY0MmZjOTAwOTAzM2IyMTJjMzNjMWFkYWY5YTliNDBiOTI3NWI4Y2Y0NzlmNjViZGNiY2I4MGY1NDJkNjA1MTEwY2M0NzI3Y2E1ZGU1NTk1ODQwOTM0Y2E5MGY0Zjg5NjQ1ZmQ4NTEwYzk4OGY3N2VhYzczOWU3Y2VlMWQ1ZTAzNWUxNmFiNWM0YzFkM2NlOTJkZGMwNDZjN2JiY2Y2MzEyOTczYjExNDc3Y2Y0Y2JlOWExMWZhOWExOWUyM2Q2ODkzNTEzNDE1MzdlOGJmMmExY2U2Mjk1OGJkNjU4YmU2MzNmZDNhN2FiNjdmZjczZTY4Nzk5YTZlNzBiNzczMGIxM2VmYWUyZWJkZDc5MDMwNDI0OWU2MTZjNmMyODY1ZDJjNzM2MmVmY2RiNTYyMzIwZDRkYTJlYzA0YTU=", "loveCode！！！@@@%%%结算小组"},
		{
			"MlPl8KhA8kt3t9l/fYxmnOQyLaTS8H1xxRrK6ZC1eF8U/CnTHy8bBoaewh9G0b9a/XXqMRlZrU7ujmUkI+0+fdVi2twKgJiRAmYUtbijLjSPMTRSk5XQFX8E0ZcJwLgirP4jeOx9cx6jA24Sc9tfJpKJUV+L+xIvSa66ZK1mrWH8crAqWkhHgtUYv9Cd12lh8YK3gItcfnauCEzxKMnMcRVlmhz5oY9zqZj7INrZE29Bdkbm2KdKxwPIJDBbMakAaaYc9mqL3WfskGsidaP/NYRC4EZEzTeJTjbJq1UhVTmpNHjRDZaPV/oZtj2zJRekHgdUcauvDiJuUpE/AWks2g==",
			"LoveCoder!结算系统小组LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!LoveCoder!",
		},
	}

	for _, testV := range testCases {
		cipherTextByte, _ := base64.StdEncoding.DecodeString(testV.cipherText)
		rawText, err := DecryptWithRSA(cipherTextByte)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, testV.rawText, string(rawText))
	}
}

func TestEncryptWithRSA(t *testing.T) {
	testCases := []struct {
		rawText    string
		cipherText string
	}{
		{"loveCode！！！@@@%%%结算小组", "MTdhOTc2NmFjZTdhODkxN2RhNmFiNjRkZjZlZTU5ZTkwYjlmMTJjMzVlZjI5MjkzN2MxMzQ5YjFjNTU0MjczYjBhY2ZhMjVhYzVmZjQ4MTU3YzY1OGQwYzVjZDYyODNiMDFkMjgwZmFiYjdjOWQxZWRhMzBkZmRjYTdiNmYxYjM3ZjQwOTRlNGNiMGI5ZDM5N2YxYzdkNmEwMjBiODAzMmFiOTAwYmZlNzYyNjUwOGM5MTQ5OGViYWJiNTYzMzM3NTEyNzY0MmZjOTAwOTAzM2IyMTJjMzNjMWFkYWY5YTliNDBiOTI3NWI4Y2Y0NzlmNjViZGNiY2I4MGY1NDJkNjA1MTEwY2M0NzI3Y2E1ZGU1NTk1ODQwOTM0Y2E5MGY0Zjg5NjQ1ZmQ4NTEwYzk4OGY3N2VhYzczOWU3Y2VlMWQ1ZTAzNWUxNmFiNWM0YzFkM2NlOTJkZGMwNDZjN2JiY2Y2MzEyOTczYjExNDc3Y2Y0Y2JlOWExMWZhOWExOWUyM2Q2ODkzNTEzNDE1MzdlOGJmMmExY2U2Mjk1OGJkNjU4YmU2MzNmZDNhN2FiNjdmZjczZTY4Nzk5YTZlNzBiNzczMGIxM2VmYWUyZWJkZDc5MDMwNDI0OWU2MTZjNmMyODY1ZDJjNzM2MmVmY2RiNTYyMzIwZDRkYTJlYzA0YTU="},
	}

	for _, testV := range testCases {
		cipherText, err := EncryptWithRSA([]byte(testV.rawText))
		if err != nil {
			t.Fatal(err)
		}
		cipherTextB64 := base64.StdEncoding.EncodeToString(cipherText)
		assert.Equal(t, testV.cipherText, cipherTextB64)
	}
}

package encryption

import (
	"fmt"
	"testing"
)

func TestGeneratePasswdPBKDF2Key(t *testing.T) {
	testCases := []struct {
		rawtext string
		saltStr string
		passwd  string
	}{
		{"123456", "9AbQYARjuk", "261bceeefd0d72f1b5cf98993798f6056b8ebfa26265dd26"},
		{"1234567a", "CFfLGNEnK3", "fa7f2b94a3503e07c4241d3bdd0c9107fbe044f5d90ea919"},
		{"aaa123", "8dqfx43L74", "0cbc5ea7e2e79fc3c2e3328c41581e227a7c53a2efbb90cc"},
	}

	for _, testV := range testCases {
		passwd, _ := DefaultPBKDF2Options.GeneratePasswdPBKDF2Key([]byte(testV.rawtext), []byte(testV.saltStr))
		fmt.Printf("[%s]用盐[%s]加密后得到：[%s]\n", testV.rawtext, testV.saltStr, passwd)
		if passwd != testV.passwd {
			t.Errorf("[%s]加密后得到的密文与预期不符\n", testV.rawtext)
		}
	}
}

func TestCheckPBKDF2PasswdForVzhuoTaskSYS(t *testing.T) {
	testCases := []struct {
		rawtext string
		saltStr string
		passwd  string
	}{
		{"123456", "9AbQYARjuk", "261bceeefd0d72f1b5cf98993798f6056b8ebfa26265dd26"},
		{"1234567a", "CFfLGNEnK3", "fa7f2b94a3503e07c4241d3bdd0c9107fbe044f5d90ea919"},
		{"aaa123", "8dqfx43L74", "0cbc5ea7e2e79fc3c2e3328c41581e227a7c53a2efbb90cc"},
	}

	for _, testV := range testCases {
		checkReslt := DefaultPBKDF2Options.CheckPBKDF2PasswdForVzhuoTaskSYS(testV.rawtext, testV.passwd, testV.saltStr)
		if checkReslt != true {
			t.Errorf("验证[%s]失败\n", testV.rawtext)
		}
	}
}

package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GenRSAkey(bits int, isReGenerateKey bool) error {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	// fmt.Println(string(pem.EncodeToMemory(priBlock)))
	// string(pem.EncodeToMemory(priBlock))

	// 生成公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	// fmt.Println(string(pem.EncodeToMemory(publicBlock)))

	codeContent := makeRSAKeyGoCode(pem.EncodeToMemory(priBlock), pem.EncodeToMemory(publicBlock))

	codeFilePath := getKeyGoCodeFilePath()

	if isExistPath, _ := PathIsExist(codeFilePath); isExistPath && !isReGenerateKey {
		return errors.New("应用程序公私钥已经生成了，目前不能重新生成")
	}

	return writeKeyGoCode(codeContent, codeFilePath)
}

// PathIsExist 判断某个路径是否存在
func PathIsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if nil == err {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func makeRSAKeyGoCode(rsaPrivateKeyByte []byte, rsaPublicKeyByte []byte) string {
	codeContent := rsaKeyGoCodeTmpl
	codeContent = strings.Replace(codeContent, "{{APPPrivateKey}}", "`"+string(rsaPrivateKeyByte)+"`", -1)
	codeContent = strings.Replace(codeContent, "{{APPPublicKey}}", "`"+string(rsaPublicKeyByte)+"`", -1)

	return codeContent
}

func getKeyGoCodeFilePath() string {
	_, b, _, _ := runtime.Caller(0)
	fpath := filepath.Dir(b)
	return path.Join(fpath, "rsa_key"+".go")
}

func writeKeyGoCode(code, fpath string) error {
	//fpath = path.Join(fpath, "rsa_key"+".go")

	f, err := os.Create(fpath)
	if err != nil {
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	_, err = f.WriteString(code)

	return nil
}

func EncryptWithRSA(raw []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(GetAPPPbulicKeyContent())
	if block == nil {
		return nil, errors.New("public key error")
	}

	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, raw)
}

func DecryptWithRSA(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(GetAPPPrivateKeyContent())
	if block == nil {
		return nil, errors.New("private key error!")
	}

	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

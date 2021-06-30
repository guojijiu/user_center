package encryption

var rsaKeyGoCodeTmpl = `package encryption

var (
	appPrivateKey = []byte({{APPPrivateKey}})
	appPublicKey  = []byte({{APPPublicKey}})
)

func GetAPPPrivateKeyContent() []byte {
	return appPrivateKey
}

func GetAPPPbulicKeyContent() []byte {
	return appPublicKey
}
`

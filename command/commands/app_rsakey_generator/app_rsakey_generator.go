package app_rsakey_generator

import (
	"fmt"
	"user_center/command"
	"user_center/pkg/color"
	"user_center/pkg/encryption"
)

var (
	keyBitCount     int = 1024
	isReGenerateKey bool
)

var CMDRSAkeyGenerator = &command.Command{
	UsageLine: "generate:rsakey",
	Short:     "user_center应用自动生成工具",
	Long:      `generate:rsakey 命令会创建一个user_center应用程序公私钥，用于加解密重要的敏感数据。`,
	Run:       createUIMSRSAKey,
}

func init() {
	CMDRSAkeyGenerator.Flag.IntVar(&keyBitCount, "b", 1024, "-b：密钥bit位数")
	CMDRSAkeyGenerator.Flag.BoolVar(&isReGenerateKey, "isReGenerateKey", false, "-isReGenerateKey：是否重新生成密钥")
	command.CMD.Register(CMDRSAkeyGenerator)
}

func createUIMSRSAKey(cmd *command.Command, args []string) int {
	if len(args) > 0 {
		err := cmd.Flag.Parse(args[1:])
		if err != nil {
			fmt.Printf("创建RSA KEY失败：%s\n", err.Error())
		}
	}

	fmt.Println("开始生成user_center app rsa key...")

	//err := exec.Command("openssl", "genrsa", "-out", "./uims_app_rsa_private_key.pem", string(keyBitCount)).Run()
	err := encryption.GenRSAkey(keyBitCount, isReGenerateKey)
	if err != nil {
		fmt.Println(color.Red(err.Error()))
	}

	fmt.Println("生成user_center app rsa key完毕！")

	return 0
}

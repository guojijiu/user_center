package app_key_generator

import (
	"fmt"
	"user_center/command"
	"user_center/pkg/color"
	"user_center/pkg/encryption"
	"user_center/pkg/env"
)

const APP_KEY_NAME = "APP_KEY"

var CMDappKeyGenerator = &command.Command{
	UsageLine: "generate:key",
	Short:     "user_permission应用自动生成工具",
	Long:      `generate:key 命令会创建一个user_permission应用程序密钥，用于加解密。`,
	Run:       createUIMSappKey,
}

func init() {
	command.CMD.Register(CMDappKeyGenerator)
}

// createUIMSappKey 创建一个新的应用密钥并写入.env文件
func createUIMSappKey(_ *command.Command, _ []string) int {
	appKey, err := encryption.GenerateBase64Key()
	if err != nil {
		fmt.Println(color.Red(err.Error()))
	}

	fmt.Println(color.Bold(appKey))

	err = env.SetKeyStringV(APP_KEY_NAME, appKey)
	if err != nil {
		fmt.Println(color.Red(err.Error()))
	}

	fmt.Println(color.Green("Generate app key successfully."))

	return 0
}

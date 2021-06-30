package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
)

var envFilePath string

func init() {
	// 嵌套查找父文件夹中的 .env 文件
	envFilePath = ".env"
	errCount := 0
	success := false
	for errCount < 10 {
		err := godotenv.Load(envFilePath)
		if err != nil {
			envFilePath = fmt.Sprintf("../%s", envFilePath)
		} else {
			success = true
			break
		}
		errCount++
	}
	if !success {
		panic("Error loading .env file")
	}
}

func DefaultGet(key string, def interface{}) interface{} {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func DefGetStr(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func DefaultGetBool(key string, def bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatal(err)
	}
	return boolVal
}

func DefaultGetInt(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	intVal, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(intVal)
}

// SetKeyStringV 设置key=value形式的配置数据写入到.env文件中
func SetKeyStringV(key, v string) error {
	bytesContent, err := ioutil.ReadFile(envFilePath)

	regexStr := "(?m)^" + key + "=.*$"
	pRegexp := regexp.MustCompile(regexStr)
	replString := []byte(key + "=" + v)
	bytesContentAfterReplace := pRegexp.ReplaceAllLiteral(bytesContent, replString)
	//fmt.Println(string(bytesContentAfterReplace))
	err = ioutil.WriteFile(envFilePath, bytesContentAfterReplace, 0644)

	return err
}

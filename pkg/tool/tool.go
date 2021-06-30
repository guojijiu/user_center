package tool

import (
	"bytes"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/sha1"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	texttmpl "text/template"
	"time"
	"user_center/pkg/color"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Dump(data interface{}) {
	bites, _ := json.MarshalIndent(data, "", "\t")
	log.Printf("%s", string(bites))
}

// return sql OffSet
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * 10
	}
	return result
}
func GenXid() string {
	uuid := xid.New()
	return uuid.String()
}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 切片进行去重
func RemoveRepByMap(slc []int) []int {
	result := []int{}         //存放返回的不重复切片
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

// 雪花算法生成19位唯一id
func GenerateUuid() string {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(1)

	// Generate a snowflake ID.
	id := node.Generate().String()
	return id
}

// 进行密码加密的方法
func EncryptPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
func StrToTime(s string) time.Time {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, s)
	return t
}

// GenerateRandStrWithMath 生成指定长度指定字符的随机字符串，注意此函数不能用于加解密相关的业务中，如果// 需要用于密码加解密相关业务中，请使用 GenerateRandStrWithCrypto
func GenerateRandStrWithMath(n int, allowedChars ...[]byte) string {
	var defaultLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var letters []byte

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// GenerateRandStrWithCrypto 生成指定长度指定字符的随机字符串，效率较 GenerateRandStrWithMath 慢
func GenerateRandStrWithCrypto(n int, allowedChars ...[]byte) string {
	return string(GenerateRandBytesWithCrypto(n, allowedChars...))
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

// IsExistPath 判断某个目录或文件是否存在
func IsExistPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

// MD5
func MD5(data string) string {
	md5Obj := md5.New()
	md5Obj.Write([]byte(data))
	return hex.EncodeToString(md5Obj.Sum(nil))
}

// SHA1
func SHA1(data string) string {
	sha1Obj := sha1.New()
	sha1Obj.Write([]byte(data))
	return hex.EncodeToString(sha1Obj.Sum(nil))
}

func BasenameNotSuffix(filename string) string {
	if len(filename) == 0 {
		return filename
	}
	suffix := filepath.Ext(filename)
	return strings.TrimSuffix(filename, suffix)
}

// PrintErrorTmplAndExit print error message by err text template.
func PrintErrorTmplAndExit(message, errTemplate string) {
	TmplTextParseAndOutput(fmt.Sprintf(errTemplate, message), nil)
	os.Exit(2)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// Endline return a new newline escape character
func EndLine() string {
	return "\n"
}

// TmplToString parses a text template and return the result as a string.
func TmplToString(tmpl string, data interface{}) string {
	t := texttmpl.New("tmpl").Funcs(UIMSfuncMap())
	texttmpl.Must(t.Parse(tmpl))

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	MustCheck(err)

	return doc.String()
}

func UIMSfuncMap() texttmpl.FuncMap {
	return texttmpl.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       color.Bold,
		"blueblod":   color.BlueBold,
		"green":      color.Green,
		"greenbold":  color.GreenBold,
		"headline":   color.MagentaBold,
		"foldername": color.RedBold,
		"endline":    EndLine,
		"tmpltostr":  TmplToString,
	}
}

func TmplTextParseAndOutput(textTmpl string, data interface{}) {
	output := color.NewColorWriter(os.Stderr)

	t := texttmpl.New("Usage").Funcs(UIMSfuncMap())
	texttmpl.Must(t.Parse(textTmpl))
	err := t.Execute(output, data)
	if err != nil {
		logrus.Error(err.Error())
	}
}

//bigint转化为时间数据
func BigIntConvertTime(intTime int) (strTime time.Time) {
	convertStrTime := strconv.Itoa(intTime)
	sliStrTime := convertStrTime[:len(convertStrTime)-6]
	convertIntTime, _ := strconv.Atoi(sliStrTime)
	//timeTemplate := "2006-01-02 15:04:05"
	//strTime = time.Unix(int64(convertIntTime), 0).Format(timeTemplate)
	strTime = time.Unix(int64(convertIntTime), 0)
	return strTime
}

// RandIntInRange 随机生成指定区间的整数
func RandIntInRange(min, max int64) int64 {
	if min > max {
		panic(errors.New("the min is greater than the max!"))
	}
	if min < 0 {
		i64Min := int64(math.Abs(float64(min)))
		result, _ := crand.Int(crand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := crand.Int(crand.Reader, big.NewInt(max-min+1))

		return min + result.Int64()
	}
}

// 将数据格式化为 JSON 字符串, 并去除末尾的 \n 符号
func JSON(v interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(v)
	if err != nil {
		return []byte{}, err
	}
	// 去除末尾的 \n
	return bf.Bytes()[:bf.Len()-1], nil
}

func JSONString(v interface{}) string {
	b, _ := JSON(v)
	return string(b)
}

//校验邮箱地址是否正确
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

package Responses

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mailru/easyjson/buffer"
	"net/http"
	validator2 "user_center/app/validator"
)

const CodeSuccess = 1
const CodeFailed = 0
const TokenInvalid = 401

func JSON(c *gin.Context, code int, message string, body interface{}) {
	c.JSON(http.StatusOK, struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Body    interface{} `json:"body"`
	}{
		Code:    code,
		Message: message,
		Body:    body,
	})
}

func Success(c *gin.Context, message string, body interface{}) {
	if message == "" {
		message = "success"
	}
	JSON(c, CodeSuccess, message, body)
}

func TokenFailed(c *gin.Context, message string) {
	if message == "" {
		message = "failed"
	}
	JSON(c, TokenInvalid, message, nil)
}

func Failed(c *gin.Context, message string, body interface{}) {
	if message == "" {
		message = "failed"
	}
	JSON(c, CodeFailed, message, body)
}

func Error(c *gin.Context, err error) {
	JSON(c, CodeFailed, err.Error(), nil)
}

// 请求参数不合法
func BadReq(c *gin.Context, err error) {
	if e, ok := err.(validator.ValidationErrors); ok {
		var errMsg buffer.Buffer
		te := e.Translate(validator2.Trans)
		// 取第一个错误信息
		for _, value := range te {
			errMsg.AppendString(value)
			break
		}
		JSON(c, CodeFailed, string(errMsg.Buf), nil)
	} else {
		Error(c, err)
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func Parse(bytes []byte) Response {
	var response Response
	_ = json.Unmarshal(bytes, &response)
	return response
}

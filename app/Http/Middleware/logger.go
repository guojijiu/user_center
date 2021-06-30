package Middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
	"time"
	"user_center/pkg/glog"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// 记录请求与响应的数据
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, ".js") ||
			strings.Contains(c.Request.URL.Path, ".css") ||
			strings.Contains(c.Request.URL.Path, ".jpg") ||
			strings.Contains(c.Request.URL.Path, ".map") ||
			strings.Contains(c.Request.URL.Path, ".png") ||
			strings.Contains(c.Request.URL.Path, "html") ||
			strings.Contains(c.Request.URL.Path, ".ico") {
			// js, css 不需要日志
			c.Next()
			return
		}
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter

		//开始时间
		startTime := time.Now()

		var request map[string]interface{}

		b, _ := c.Copy().GetRawData()

		_ = json.Unmarshal(b, &request)

		s, _ := json.Marshal(request)

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(b))

		//处理请求
		c.Next()

		responseBody := bodyLogWriter.body.String()

		//结束时间
		endTime := time.Now()

		glog.Channel("request").WithFields(log.Fields{
			"request_uri":    c.Request.RequestURI,
			"request_method": c.Request.Method,
			"client_ip":      c.ClientIP(),
			"request_time":   startTime.String(),
			"response_time":  endTime.String(),
			"requests":       string(s),
			"response":       responseBody,
			"use_time":       endTime.Sub(startTime).String(),
		}).Info("记录请求")
	}
}

package logr

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CtxRequestMeta struct {
	Token          string
	Method         string
	Path           string
	QueryString    string
	Body           []byte
	AppVersion     string
	OsVersion      string
	DeviceInfo     string
	AcceptLanguage string
	XIndoChatKey   string
}

func ExtractReqMeta(c *gin.Context) *CtxRequestMeta {
	meta := CtxRequestMeta{}
	meta.Token = c.GetHeader("Bearer")
	meta.AppVersion = c.GetHeader("App-Version")
	meta.OsVersion = c.GetHeader("OS-Version")
	meta.DeviceInfo = c.GetHeader("Device-Info")
	meta.AcceptLanguage = c.GetHeader("Accept-Language")
	meta.XIndoChatKey = c.GetHeader("X-IndoChat-Key")
	// when you read the the body buffer, the buffer will gone
	bodyBytes, _ := c.GetRawData()
	meta.Body = bodyBytes

	// restore the io.ReadCloser
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	query := c.Request.URL.Query()

	// handle the query string map to one row string
	b := new(bytes.Buffer)

	for qsK, q := range query {
		for _, qsV := range q {
			_, _ = fmt.Fprintf(b, "%s=\"%s\" ", qsK, qsV)
		}
	}
	// handle done...

	meta.Method = c.Request.Method
	meta.Path = c.Request.URL.EscapedPath()
	meta.QueryString = b.String()

	return &meta
}

func NewGinCtxLogger(c *gin.Context) *zap.Logger {
	logger := NewLogger("")

	// TBD body is gone
	// when you read the the body buffer, the buffer will gone
	//bodyBytes, _ := c.GetRawData()

	// restore the io.ReadCloser
	//c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	query := c.Request.URL.Query()

	// handle the query string map to one row string
	b := new(bytes.Buffer)
	for qsK, q := range query {
		for _, qsV := range q {
			_, _ = fmt.Fprintf(b, "%s=\"%s\" ", qsK, qsV)
		}
	}
	// handle done...

	method := c.Request.Method
	URI := c.Request.URL.EscapedPath()
	qs := b.String()

	logger = logger.With(
		zap.String("method", method),
		zap.String("Path", URI),
		zap.String("qs", qs),
	)

	return logger
}

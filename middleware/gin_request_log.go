package middleware

import (
	"encoding/json"

	"api-automation-backend/pkg/logr"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LogRequest a gin middleware to log every http request
func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		meta := logr.ExtractReqMeta(c)

		var rawBody interface{}
		if err := json.Unmarshal(meta.Body, &rawBody); err != nil {
			rawBody = string(meta.Body)
		}

		logr.L.Info("[Request]:",
			zap.String("token", meta.Token),
			zap.String("App-Version", meta.AppVersion),
			zap.String("OS-Version", meta.OsVersion),
			zap.String("Device-Info", meta.DeviceInfo),
			zap.String("Accept-Language", meta.AcceptLanguage),
			zap.String("X-IndoChat-Key", meta.XIndoChatKey),
			zap.String("type", "API_REQ"),
			zap.String("method", meta.Method),
			zap.String("path", meta.Path),
			zap.Any("body", rawBody),
			zap.String("qs", meta.QueryString),
		)

		c.Next()
	}
}

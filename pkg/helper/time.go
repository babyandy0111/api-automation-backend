package helper

import (
	"time"
)

// 時間格式範例 2012-01-01T12:00:00Z
func Time2ApiresTime(str time.Time) string {
	return str.Format(time.RFC3339)
}

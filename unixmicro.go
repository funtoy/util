package utils

import (
	"time"
)

func UnixMs() int64 {
	return time.Now().UnixNano() / 1000000
}

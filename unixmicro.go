package util

import (
	"time"
)

func UnixMs() int64 {
	return time.Now().UnixNano() / 1000000
}

func DayStartUnix() int64 {
	Y, M, D := time.Now().Date()
	return time.Date(Y, M, D, 0, 0, 0, 0, time.Local).Unix()
}

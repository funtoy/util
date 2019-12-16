package utils

import (
	"fmt"
	"strconv"
)

// string ToFloat64 Float64
func ToFloat64(in string) (float64, error) {
	if in == "" {
		return 0, nil
	}
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return 0, err
	}
	return out, nil
}

// Str Str
func Str(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

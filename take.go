package util

import (
	"math/rand"
)

func RandTake(n int, src []int32, r *rand.Rand) (val, left []int32) {
	left = src
	if len(left) < n || r == nil {
		return
	}
	for c := 0; c < n; c++ {
		index := r.Intn(len(left))
		val = append(val, left[index])
		left = append(left[:index], left[index+1:]...)
	}
	return
}

func Take(src []int32, t ...int32) (left []int32, ok bool) {
	left = src
	for _, v := range t {
		if !IsMemberInt32(v, left) {
			return left, false
		}
		left = DelSliceInt32(left, v)
	}
	return left, true
}

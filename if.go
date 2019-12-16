package utils

func IF(assert bool, a, b interface{}) interface{} {
	if assert {
		return a
	}
	return b
}

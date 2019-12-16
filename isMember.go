package utils

func IsMemberInt(element int, list []int) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

func IsMemberInt32(element int32, list []int32) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

func IsMemberInt64(element int64, list []int64) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

func IsMemberFloat64(element float64, list []float64) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

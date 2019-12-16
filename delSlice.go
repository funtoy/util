package utils

// DelSliceInt DelSliceInt
func DelSliceInt(list []int, i int) []int {
	for pos, v := range list {
		if v == i {
			list = append(list[:pos], list[pos+1:]...)
			return list
		}
	}
	return list
}

// DelSliceInt32 DelSliceInt32
func DelSliceInt32(list []int32, i int32) []int32 {
	for pos, v := range list {
		if v == i {
			list = append(list[:pos], list[pos+1:]...)
			return list
		}
	}
	return list
}

// DelSliceInt64 DelSliceInt64
func DelSliceInt64(list []int64, i int64) []int64 {
	for pos, v := range list {
		if v == i {
			list = append(list[:pos], list[pos+1:]...)
			return list
		}
	}
	return list
}

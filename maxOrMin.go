package util

func MaxInt64(i ...int64) int64 {
	var x int64
	for _, v := range i {
		if v > x {
			x = v
		}
	}
	return x
}

func MaxInt32(i ...int32) int32 {
	var x int32
	for _, v := range i {
		if v > x {
			x = v
		}
	}
	return x
}

func MaxInt(i ...int) int {
	var x int
	for _, v := range i {
		if v > x {
			x = v
		}
	}
	return x
}

func MinInt64(i ...int64) int64 {
	var x int64
	for _, v := range i {
		if x == 0 {
			x = v
			continue
		}
		if v < x {
			x = v
		}
	}
	return x
}

func MinInt32(i ...int32) int32 {
	var x int32
	for _, v := range i {
		if x == 0 {
			x = v
			continue
		}
		if v < x {
			x = v
		}
	}
	return x
}

func MinInt(i ...int) int {
	var x int
	for _, v := range i {
		if x == 0 {
			x = v
			continue
		}
		if v < x {
			x = v
		}
	}
	return x
}

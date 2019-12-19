package util

// 桌子位置管理器

type Seat []string

func NewSeat(v int32) Seat {
	return make(Seat, v)
}

func (s Seat) Len() int32 {
	return int32(len(s))
}

func (s Seat) Clean() {
	for i := range s {
		s[i] = ""
	}
}

func (s Seat) Find(key string) (position int32, ok bool) {
	for index, val := range s {
		if key == val {
			position = int32(index+1)
			ok = true
			return
		}
	}
	for index, val := range s {
		if val == "" {
			s[index] = key
			position = int32(index+1)
			ok = true
			return
		}
	}
	return
}

func (s Seat) Leave(key string) (hasSeat bool) {
	for index, val := range s {
		if val == key {
			s[index] = ""
			return true
		}
	}
	return false
}

func (s Seat) WhoIs(pos int32) (key string, ok bool) {
	if s.Len() < pos {
		return
	}
	for index, val := range s {
		if pos == int32(index+1) {
			key = val
			ok = true
			return
		}
	}
	return
}

func (s Seat) WhereIs(key string) (pos int32, ok bool) {
	for index, val := range s {
		if val == key {
			pos = int32(index+1)
			ok = true
			return
		}
	}
	return
}

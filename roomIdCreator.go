package util

import (
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var DefaultIdPool = NewRoomIdPool()

type roomIdPool struct {
	mu     sync.Mutex
	ids    map[string]bool
	random *rand.Rand
}

func NewRoomIdPool() *roomIdPool {
	return &roomIdPool{
		ids:    make(map[string]bool),
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (i *roomIdPool) New() string {
	i.mu.Lock()
	defer i.mu.Unlock()

	for n := 0; n < 5; n++ {
		id := i.createId()
		if _, ok := i.ids[id]; ok {
			continue
		}
		i.ids[id] = true
		return id
	}
	return "0"
}

func (i *roomIdPool) createId() string {
	var reply int32
	for n := 0; n < 6; n++ {
		x := i.random.Intn(9) + 1
		reply += int32(math.Pow(10, float64(n))) * int32(x)
	}
	return strconv.Itoa(int(reply))
}

func (i *roomIdPool) Release(id string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.ids, id)
}

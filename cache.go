package util

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var MemCache = &cache{sm: make(map[string][]byte), mu: new(sync.Mutex)}

type cache struct {
	sm       map[string][]byte
	mu       *sync.Mutex
	filename string
}

func (c *cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.sm[key]; ok {
		return v, ok
	}
	return nil, false
}

func (c *cache) Set(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.sm[key] = value
	c.save()
}

func (c *cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.sm, key)
	c.save()
}

func (c *cache) Open(filename string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.filename = filename
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.sm)
	if err != nil {
		return err
	}

	return nil
}

func (c *cache) save() error {
	if c.filename == "" {
		return nil
	}

	data, err := json.Marshal(&c.sm)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.filename, data, 0666)
}

package utils

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache struct {
	c *cache.Cache
}

func NewInMemoryCache() *Cache {
	return &Cache{c: cache.New(20*time.Minute, 30*time.Minute)}
}

// set or add a value to the cache
func (c *Cache) Set(key string, value interface{}) {
	fmt.Println("============ set cache ==============")
	fmt.Println(key, value)
	fmt.Println("============ set cache ==============")

	c.c.Set(key, value, cache.DefaultExpiration)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.c.Get(key)
}

func (c *Cache) Del(key string) {
	c.c.Delete(key)
}

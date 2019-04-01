package service

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheManager struct {
	//saving token in cache
	cacheToken *cache.Cache
	//saving other data in cache
	cacheJson *cache.Cache
}

func InitCache() *CacheManager {
	cToken := cache.New(5*time.Minute, 30*time.Minute)
	cJson := cache.New(5*time.Minute, 30*time.Minute)
	return &CacheManager{cacheToken: cToken, cacheJson: cJson}
}

//use this function for caching data
func (c *CacheManager) Cacheable(key string, callback func() interface{}) interface{} {
	data := c.CacheGet(key)
	if data == nil {
		return callback()
	} else {
		return data
	}
}

//use this function for get data
func (c *CacheManager) CacheGet(key string) interface{} {
	if x, found := c.cacheJson.Get(key); found {
		return x
	}

	return nil
}

//use this function for clearing data
func (c *CacheManager) CacheEvict(key string) {
	c.cacheJson.Delete(key)
}

//use this function for save token after created
func (c *CacheManager) SaveToken(key string, data string) {
	c.cacheToken.Set(key, data, 24*time.Hour)
}

//use this function for save token after created if not expired
func (c *CacheManager) SaveTokenNoExpired(key string, data string) {
	c.cacheToken.Set(key, data, cache.NoExpiration)
}

//use this function for get token
func (c *CacheManager) GetToken(key string, data string) interface{} {
	if x, found := c.cacheToken.Get(key); found {
		return x.(string)
	}

	return nil
}

//use this function for delete token
func (c *CacheManager) DeleteToken(key string) {
	c.cacheToken.Delete(key)
}

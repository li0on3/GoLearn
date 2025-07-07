package day03

import "time"

type Cache struct {
	data map[string]interface{}
	TTL  map[string]time.Duration
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.data[key] = value
	c.TTL[key] = ttl
}

func (c *Cache) GetValue(key string) (interface{}, bool) {
	v, ok := c.data[key]
	return v, ok
}

func (c *Cache) GetTTL(key string) time.Duration {
	v, ok := c.TTL[key]
	if !ok {
		return time.Duration(0)
	}
	return v
}

// Delete 删除单个键值    安全操作不会 panic
func (c *Cache) Delete(key string) {
	delete(c.data, key)
	delete(c.TTL, key)
}

// Clear 清空所有的缓存   直接新建两个nil的map  覆盖即可实现
func (c *Cache) Clear() {
	c.data = make(map[string]interface{})
	c.TTL = make(map[string]time.Duration)
}

package main

type Cache interface {
	Get(cid string) (string, error)
	Set(cid string)
	Delete(cid string) error
}

type Memcache struct {
}

type Redis struct {
}

type RedisCache struct { // RedisCache 继承Cache
	Cache // 只要继承interface 接口，必须实现其全部方法
}

type RedisCache2 struct { // RedisCache2 继承RedisCache
	RedisCache // 继承struct 结构体，不需要实现其全部方法
}

func (r *Redis) Get(cid string) (string, error) {
	panic("implement me")
}

func (r *Redis) Set(cid string) {
	panic("implement me")
}

func (r *Redis) Delete(cid string) error {
	panic("implement me")
}

func (m *Memcache) Get(cid string) (string, error) {
	panic("implement me")
}

func (m *Memcache) Set(cid string) {
	panic("implement me")
}

func (m *Memcache) Delete(cid string) error {
	panic("implement me")
}

func (r *RedisCache) Set(cid string) {
	panic("implement me")
}

func (r *RedisCache) Delete(cid string) error {
	panic("implement me")
}

func (m *RedisCache) Get(cid string) (string, error) {
	panic("implement me")
}

func (r *RedisCache) add(cid string) error { // RedisCache 增加 add方法
	panic("implement me")
}

func (r *RedisCache2) IsCreate(b bool) error { // RedisCache2 增加 IsCreate方法
	panic("implement me")
}

func main() {
	var c Cache

	// Redis 实现Cache的Get、Set、Delete 方法，所以Redis 是Cache的多态体现
	c = &Redis{}
	c.Get("")
	c.Set("")
	c.Delete("")

	// Memcache 实现Cache的Get、Set、Delete 方法，所以Memcache 是Cache的多态体现
	c = &Memcache{}
	c.Get("")
	c.Set("")
	c.Delete("")

	r2 := &RedisCache2{}
	r2.Get("")        // 继承Redis
	r2.Set("")        // 继承Redis
	r2.Delete("")     // 继承Redis
	r2.add("")        // 继承Redis
	r2.IsCreate(true) // RedisCache2 增加 IsCreate方法

}

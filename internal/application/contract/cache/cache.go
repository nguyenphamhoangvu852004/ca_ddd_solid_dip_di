package cache

type Cache interface {
	Get(key string) interface{}
	Save(key string, value interface{}, expiration int)
}

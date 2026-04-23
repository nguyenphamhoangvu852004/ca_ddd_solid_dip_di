package cache

import "ca_ddd_solid_dip_di/internal/application/contract/cache"

type RedisCacheImpl struct {
	redisClient interface{} // Placeholder for Redis client
}

// Get implements [cache.Cache].
func (r *RedisCacheImpl) Get(key string) interface{} {
	panic("unimplemented")
}

// Save implements [cache.Cache].
func (r *RedisCacheImpl) Save(key string, value interface{}, expiration int) {
	panic("unimplemented")
}

func NewRedisCacheImpl() cache.Cache {
	return &RedisCacheImpl{
		redisClient: nil, // Initialize Redis client here
	}
}

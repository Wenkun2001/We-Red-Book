package repository

import (
	"context"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/repository/cache"
)

var ErrCodeVerifyTooMany = cache.ErrCodeVerifyTooMany

type CodeRepository interface {
	Set(ctx context.Context, biz, phone, code string) error
	Verify(ctx context.Context, biz, phone, code string) (bool, error)
}

type CacheCodeRepository struct {
	cache *cache.RedisCodeCache
}

func NewCodeRepository(c *cache.RedisCodeCache) *CacheCodeRepository {
	return &CacheCodeRepository{
		cache: c,
	}
}

func (c *CacheCodeRepository) Set(ctx context.Context, biz, phone, code string) error {
	return c.cache.Set(ctx, biz, phone, code)
}

func (c *CacheCodeRepository) Verify(ctx context.Context, biz, phone, code string) (bool, error) {
	return c.cache.Verify(ctx, biz, phone, code)
}

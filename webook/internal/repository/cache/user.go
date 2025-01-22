package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Wenkun2001/We-Red-Book/webook/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrKeyNotExist = redis.Nil

type UserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func (c *UserCache) Get(ctx context.Context, uid int64) (domain.User, error) {
	key := c.key(uid)
	// 我假定这个地方用 JSON 来
	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal([]byte(data), &u)
	//if err != nil {
	//	return domain.User{}, err
	//}
	//return u, nil
	return u, err
}

func (c *UserCache) Set(ctx context.Context, du domain.User) error {
	key := c.key(du.Id)
	// 我假定这个地方用 JSON
	data, err := json.Marshal(du)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}

func (c *UserCache) key(uid int64) string {
	// user-info-
	// user.info.
	// user/info/
	// user_info_
	return fmt.Sprintf("user:info:%d", uid)
}

type UserCacheV1 struct {
	client *redis.Client
}

func NewUserCache(cmd redis.Cmdable) *UserCache {
	return &UserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

// 一定不要自己去初始化你需要的东西，让外面传进来
//func NewUserCacheV1(addr string) *UserCache {
//	cmd := redis.NewClient(&redis.Options{Addr: addr})
//	return &UserCache{
//		cmd:        cmd,
//		expiration: time.Minute * 15,
//	}
//}

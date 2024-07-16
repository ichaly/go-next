package base

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/store"
	sb "github.com/eko/gocache/store/bigcache/v4"
	sr "github.com/eko/gocache/store/redis/v4"
	"github.com/ichaly/go-next/lib/base/internal"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func NewStorage(v *viper.Viper) (*cache.Cache[string], error) {
	cfg := &internal.CacheConfig{}
	if err := v.Sub("cache").Unmarshal(cfg); err != nil {
		return nil, err
	}

	var s store.StoreInterface
	if strings.ToLower(cfg.Dialect) == "redis" {
		args := []interface{}{cfg.Host, cfg.Port}
		s = sr.NewRedis(redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", args...),
			Username: cfg.Username,
			Password: cfg.Password,
		}))
	} else {
		client, err := bigcache.New(context.Background(), bigcache.DefaultConfig(5*time.Minute))
		if err != nil {
			return nil, err
		}
		s = sb.NewBigcache(client)
	}
	return cache.New[string](s), nil
}

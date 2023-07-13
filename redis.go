package redis

import (
	"context"
	"fmt"

	"github.com/AleksandrMac/redis/cluster"
	"github.com/AleksandrMac/redis/sentinel"
	"github.com/AleksandrMac/redis/single"
	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Push(val []byte) error
	Pop() ([]byte, error)
}

type Options struct {
	Addrs              []string
	Username           string
	Password           string
	DB                 int
	SentinelMasterName string
}

func NewRedisSingleClient(ctx context.Context, opt *Options) (RedisClient, error) {
	switch len(opt.Addrs) {
	case 0:
		return nil, fmt.Errorf("redis addres not found")
	case 1:
	default:
		return nil, fmt.Errorf("redis single client does not support multiple address")
	}

	return single.New(context.Background(), &redis.Options{
		Addr:     opt.Addrs[0],
		Username: opt.Username,
		Password: opt.Password,
		DB:       opt.DB,
	})
}

func NewRedisClusterClient(ctx context.Context, opt *Options) (RedisClient, error) {
	if len(opt.Addrs) == 0 {
		return nil, fmt.Errorf("redis addres not found")

	}

	return cluster.New(ctx, &redis.ClusterOptions{
		Addrs:    opt.Addrs,
		Username: opt.Username,
		Password: opt.Password,
	})
}

func NewRedisSentinalClient(ctx context.Context, opt *Options) (RedisClient, error) {
	switch len(opt.Addrs) {
	case 0:
		return nil, fmt.Errorf("redis addres not found")
	case 1:
	default:
		return nil, fmt.Errorf("redis single client does not support multiple address")
	}

	return sentinel.New(ctx, &redis.FailoverOptions{
		MasterName:    opt.SentinelMasterName,
		SentinelAddrs: opt.Addrs,
		Username:      opt.Username,
		Password:      opt.Password,
		DB:            opt.DB,
	})
}

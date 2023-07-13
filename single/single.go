package single

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisSingle struct {
	client *redis.Client
}

func New(ctx context.Context, opt *redis.Options) (*RedisSingle, error) {
	RS := RedisSingle{}
	RS.client = redis.NewClient(opt)
	return &RS, RS.client.Ping(ctx).Err()
}

func (x *RedisSingle) Push(val []byte) error {
	return nil
}

func (x *RedisSingle) Pop() ([]byte, error) {
	return nil, nil
}

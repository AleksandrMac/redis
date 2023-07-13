package sentinel

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisSentinel struct {
	client *redis.Client
}

func New(ctx context.Context, opt *redis.FailoverOptions) (*RedisSentinel, error) {
	RS := RedisSentinel{
		client: redis.NewFailoverClient(opt),
	}
	return &RS, RS.client.Ping(ctx).Err()
}

func (x *RedisSentinel) Push(val []byte) error {
	x.client.
	return nil
}

func (x *RedisSentinel) Pop() ([]byte, error) {
	return nil, nil
}

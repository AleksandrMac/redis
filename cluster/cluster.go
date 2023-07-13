package cluster

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisCluster struct {
	client *redis.ClusterClient
}

func New(ctx context.Context, opt *redis.ClusterOptions) (*RedisCluster, error) {
	RC := RedisCluster{
		client: redis.NewClusterClient(opt),
	}

	return &RC, RC.client.Ping(ctx).Err()
}

func (x *RedisCluster) Push(val []byte) error {
	return nil
}

func (x *RedisCluster) Pop() ([]byte, error) {
	return nil, nil
}

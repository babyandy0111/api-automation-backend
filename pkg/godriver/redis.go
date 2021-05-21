package godriver

import (
	"github.com/go-redis/redis/v7"
	"log"
	"sync"
)

var (
	rc     *redis.ClusterClient
	rcOnce sync.Once
)

// NewRedis return singleton redis cluster instance
// base on go-redis v7
func NewRedis(addr []string) (*redis.ClusterClient, error) {
	var err error
	rcOnce.Do(func() {
		rc, err = newRedis(addr)
	})
	return rc, err
}

func newRedis(addr []string) (*redis.ClusterClient, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{Addrs: addr})

	if _, err := client.Ping().Result(); err != nil {
		log.Println("redis err", err)
		return nil, err
	}

	return client, nil
}

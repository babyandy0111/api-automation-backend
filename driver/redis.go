package driver

import (
	"api-automation-backend/config"
	"fmt"
	"os"
	"strconv"

	"github.com/go-redis/redis/v7"
	"github.com/indochat/godriver"
)

const RedisMasterCount = 3
const RedisSlavePerMasterCount = 1

func NewRedis() (*redis.ClusterClient, error) {
	var addr []string

	if config.GetEnvironment() != config.EnvLocalhost {
		addr = []string{os.Getenv("REDIS_CLUSTER_CONFIGURATION_POINT")}
	} else {
		// 本機開發用
		nodeCount := RedisMasterCount * (RedisSlavePerMasterCount + 1)
		addr = make([]string, nodeCount)

		for i := 0; i <= nodeCount-1; i++ {
			env := fmt.Sprintf("REDIS_CLUSTER_%s", strconv.Itoa(i))
			node := os.Getenv(env)
			addr[i] = node
		}
	}

	return godriver.NewRedis(addr)
}

func NewSingleRedis() (*redis.Client, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client, nil
}

package redis

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/xiaobudongzhang/micro-basic/config"
)

var (
	client *redis.Client
	m      sync.RWMutex
	inited bool
)

func Init() {

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("已经初始化过redis....")
		return
	}

	redisConfig := config.GetRedisConfig()
	fmt.Println(redisConfig.GetEnabled())
	if redisConfig != nil && redisConfig.GetEnabled() {
		log.Log("初始化redis...")

		if redisConfig.GetSentinelConfig() != nil && redisConfig.GetSentinelConfig().GetEnabled() {
			log.Log("初始化哨兵模式...")
			initSentinel(redisConfig)
		} else {
			log.Log("初始化普通模式...")
			initSingle(redisConfig)
		}

		log.Log("初始化redis 检测连接...")

		pong, err := client.Ping().Result()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Log("ping %s", pong)
	}

	inited = true
}

func GetRedis() *redis.Client {
	return client
}

func initSentinel(redisConfig config.RedisConfig) {
	client = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    redisConfig.GetSentinelConfig().GetMaster(),
		SentinelAddrs: redisConfig.GetSentinelConfig().GetNodes(),
		DB:            redisConfig.GetDBNum(),
		Password:      redisConfig.GetPassword(),
	})
}

func initSingle(redisConfig config.RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.GetConn(),
		Password: redisConfig.GetPassword(),
		DB:       redisConfig.GetDBNum(),
	})
}

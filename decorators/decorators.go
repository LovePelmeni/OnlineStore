package decorators

import (
	"fmt"
	"os"
	"github.com/go-redis/redis"
	"time"
	"logging"
	"errors"
)


var (
	DebugLogger = logging.DebugLogger 
	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")

	redis_client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", RedisHost, RedisPort),
		Password: fmt.Sprintf("%s", RedisPassword),
		DB: 0,
	})
)

func SetCachedValue(key string, value interface{}, expires_at time.Duration) (bool, interface{}){
	error_value := redis_client.Set(key, value, expires_at).Err().Error()
	DebugLogger.Println(fmt.Sprintf(
	"Data with Key: '%s' has been set to redis storage. Will Expire At. %s", key, expires_at))
	switch error_value{
		case error_value:
			return false, error_value
	}	
	return true, nil 
}




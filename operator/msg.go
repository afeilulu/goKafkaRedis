package operator

import (
	"encoding/json"

	"afeilulu.com/goKafkaRedis/config"
	"afeilulu.com/goKafkaRedis/model"
	"github.com/go-redis/redis"
)

func Handle(msgBytes []byte) {
	var msgObj model.Msg
	json.Unmarshal(msgBytes, &msgObj)

	if len(msgObj.Http.Http_user_agent) > 0 {
		// count signature by ip through hyperloglog
		config.RedisClient.PFAdd(msgObj.Alert.Signature, msgObj.Src_ip)

		config.RedisClient.PFCount(msgObj.Alert.Signature).Val()
		z := new(redis.Z)
		z.Score = float64(config.RedisClient.PFCount(msgObj.Alert.Signature).Val())
		z.Member = msgObj.Alert.Signature
		config.RedisClient.ZAdd("signature", *z)
	}
}

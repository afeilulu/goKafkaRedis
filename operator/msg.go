package operator

import (
	"encoding/json"
	"strconv"

	"afeilulu.com/goKafkaRedis/config"
	"afeilulu.com/goKafkaRedis/model"
	"github.com/go-redis/redis"
)

func Handle(msgBytes []byte) {
	var msgObj model.Msg
	json.Unmarshal(msgBytes, &msgObj)

	if len(msgObj.Http.Http_user_agent) > 0 {
		key := msgObj.Alert.Signature + " " + strconv.FormatInt(msgObj.Alert.Signature_id,10)
		// count signature by ip through hyperloglog
		config.RedisClient.PFAdd(key, msgObj.Src_ip)

		z := new(redis.Z)
		z.Score = float64(config.RedisClient.PFCount(key).Val())
		z.Member = key
		config.RedisClient.ZAdd("signature", *z)
	}
}

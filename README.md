# goKafkaRedis
Consume message from kafka, do some statistics work and save into redis.
Input message json format.
```
{"timestamp":"","src_ip":"192.168.0.1","http":{"http_user_agent":"value"},"alert":{"signature":"feature"}}
```

# build
```
$go build
```

## run
```
./goKafkaRedis broker group topic
```

## show result
reverse show top 100 in redis zsort
```
$redis-cli -h host -p port -a password ZREVRANGE signature 0 100 withscores
```
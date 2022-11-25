package dao

import "github.com/go-redis/redis"

var Rdb *redis.Client

func Initredis() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err = Rdb.Ping().Result()
	return err
}

func Upin(username string) {
	err := Rdb.Set(username, 1, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Upout(username string) (flag int) {
	_, err := Rdb.Get(username).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		panic(err)
		return
	}
	return 1
}

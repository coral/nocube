package data

import (
	"github.com/go-redis/redis"
)

type Data struct {
	db *redis.Client
}

func New() Data {
	return Data{}
}

func (d *Data) Init() {
	d.db = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := d.db.Ping().Result()
	if err != nil {
		panic("No redis")
	}

}

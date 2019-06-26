package data

import (
	"strconv"

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

////FLOAT64

func (d *Data) SetFloat64(name string, key string, value float64) {
	err := d.db.Set(name+"_"+key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (d *Data) GetFloat64(name string, key string) float64 {
	val, err := d.db.Get(name + "_" + key).Result()
	if err != nil {
		return 0.0
	}
	f, _ := strconv.ParseFloat(val, 64)
	return f

}

//Int64
func (d *Data) SetInt64(name string, key string, value int64) {
	err := d.db.Set(name+"_"+key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (d *Data) GetInt64(name string, key string) int64 {
	val, err := d.db.Get(name + "_" + key).Result()
	if err != nil {
		return 0
	}
	f, _ := strconv.ParseInt(val, 10, 64)
	return f

}

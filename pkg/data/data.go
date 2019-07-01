package data

import (
	"strconv"

	"github.com/go-redis/redis"
)

type Data struct {
	db         *redis.Client
	floatcache map[string]float64
	intcache   map[string]int64
}

func New() Data {
	return Data{
		floatcache: make(map[string]float64),
		intcache:   make(map[string]int64),
	}
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

func (d *Data) SetScopedFloat64(pipeline string, effect string, key string, value float64) {
	err := d.db.Set(pipeline+"_"+effect+"_"+key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	d.floatcache[pipeline+"_"+effect+"_"+key] = value
}

func (d *Data) GetScopedFloat64(pipeline string, effect string, key string) float64 {
	if d, m := d.floatcache[pipeline+"_"+effect+"_"+key]; m {
		return d
	}
	val, err := d.db.Get(pipeline + "_" + effect + "_" + key).Result()
	if err != nil {
		return 0.0
	}
	f, _ := strconv.ParseFloat(val, 64)
	d.floatcache[pipeline+"_"+effect+"_"+key] = f
	return f

}

//Int64
func (d *Data) SetScopedInt64(pipeline string, effect string, key string, value int64) {
	err := d.db.Set(pipeline+"_"+effect+"_"+key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	d.intcache[pipeline+"_"+effect+"_"+key] = value
}

func (d *Data) GetScopedInt64(pipeline string, effect string, key string) int64 {
	if d, m := d.intcache[pipeline+"_"+effect+"_"+key]; m {
		return d
	}
	val, err := d.db.Get(pipeline + "_" + effect + "_" + key).Result()
	if err != nil {
		return 0
	}
	f, _ := strconv.ParseInt(val, 10, 64)
	d.intcache[pipeline+"_"+effect+"_"+key] = f
	return f

}

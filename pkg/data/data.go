package data

import (
	"strconv"
	"sync"

	"github.com/go-redis/redis"
)

type Data struct {
	db         *redis.Client
	floatcache map[string]float64
	floatMutex sync.RWMutex
	intcache   map[string]int64
	intMutex   sync.RWMutex
}

func New() Data {
	return Data{
		floatcache: make(map[string]float64),
		floatMutex: sync.RWMutex{},
		intcache:   make(map[string]int64),
		intMutex:   sync.RWMutex{},
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
	d.floatMutex.Lock()
	d.floatcache[pipeline+"_"+effect+"_"+key] = value
	d.floatMutex.Unlock()
}

func (d *Data) GetScopedFloat64(pipeline string, effect string, key string) float64 {
	d.floatMutex.RLock()
	if val, m := d.floatcache[pipeline+"_"+effect+"_"+key]; m {
		d.floatMutex.RUnlock()
		return val
	}
	d.floatMutex.RUnlock()
	val, err := d.db.Get(pipeline + "_" + effect + "_" + key).Result()
	if err != nil {
		return 0.0
	}
	f, _ := strconv.ParseFloat(val, 64)
	d.floatMutex.Lock()
	d.floatcache[pipeline+"_"+effect+"_"+key] = f
	d.floatMutex.Unlock()
	return f

}

//Int64
func (d *Data) SetScopedInt64(pipeline string, effect string, key string, value int64) {
	err := d.db.Set(pipeline+"_"+effect+"_"+key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	d.intMutex.Lock()
	d.intcache[pipeline+"_"+effect+"_"+key] = value
	d.intMutex.Unlock()
}

func (d *Data) GetScopedInt64(pipeline string, effect string, key string) int64 {
	d.intMutex.RLock()
	if val, m := d.intcache[pipeline+"_"+effect+"_"+key]; m {
		d.intMutex.RUnlock()
		return val
	}
	d.intMutex.RUnlock()
	val, err := d.db.Get(pipeline + "_" + effect + "_" + key).Result()
	if err != nil {
		return 0
	}
	f, _ := strconv.ParseInt(val, 10, 64)
	d.intMutex.Lock()
	d.intcache[pipeline+"_"+effect+"_"+key] = f
	d.intMutex.Unlock()
	return f

}

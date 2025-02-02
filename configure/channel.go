package configure

import (
	"fmt"
	"strconv"

	"github.com/gwuhaolin/livego/utils/uid"

	"github.com/go-redis/redis/v7"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type RoomKeysType struct {
	redisCli   *redis.Client
	localCache *cache.Cache
}

var RoomKeys = &RoomKeysType{
	localCache: cache.New(cache.NoExpiration, 0),
}

var saveInLocal = true

func Init() {
	saveInLocal = len(Config.GetString("redis_addr")) == 0
	if saveInLocal {
		return
	}

	RoomKeys.redisCli = redis.NewClient(&redis.Options{
		Addr:     Config.GetString("redis_addr"),
		Password: Config.GetString("redis_pwd"),
		DB:       0,
	})

	_, err := RoomKeys.redisCli.Ping().Result()
	if err != nil {
		log.Panic("Redis: ", err)
	}

	log.Info("Redis connected")
}

func (r *RoomKeysType) AddViewer(channel string) (err error) {
	viewers, err := r.redisCli.Get(channel + ":stream").Result()
	if err != nil {
		log.Warn("[VIEWERS] ", err)
		return err
	}
	viewerNum, err := strconv.Atoi(viewers)
	if err != nil {
		log.Warn("[VIEWERS] ", err)
		return err
	}
	r.redisCli.Set(channel+":stream", viewerNum+1, 0)
	return
}

func (r *RoomKeysType) SubtractViewer(channel string) (err error) {
	log.Debug(channel)
	viewers, err := r.redisCli.Get(channel + ":stream").Result()
	if err != nil {
		log.Warn("[VIEWERS] ", err)
		return
	}
	viewerNum, err := strconv.Atoi(viewers)
	if err != nil {
		log.Warn("[VIEWERS] ", err)
		return
	}
	r.redisCli.Set(channel+":stream", viewerNum-1, 0)
	return
}

func (r *RoomKeysType) DeleteStream(channel string) (err error) {
	log.Debug("[STREAM] delete ", channel)
	r.redisCli.Del(channel + ":stream")
	return
}

func (r *RoomKeysType) SetStream(channel string) (err error) {
	r.redisCli.Set(channel+":stream", -1, 0)
	return
}

// set/reset a random key for channel
func (r *RoomKeysType) SetKey(channel string) (key string, err error) {
	if !saveInLocal {
		for {
			key = uid.RandStringRunes(48)
			if _, err = r.redisCli.Get(key).Result(); err == redis.Nil {
				err = r.redisCli.Set(channel+":key", key, 0).Err()
				if err != nil {
					return
				}
				err = r.redisCli.Set("key:"+key, channel, 0).Err()
				return
			} else if err != nil {
				return
			}
		}
	}

	for {
		key = uid.RandStringRunes(48)
		if _, found := r.localCache.Get(key); !found {
			r.localCache.SetDefault(channel, key)
			r.localCache.SetDefault(key, channel)
			break
		}
	}
	return
}

func (r *RoomKeysType) GetKey(channel string) (newKey string, err error) {
	if !saveInLocal {
		if newKey, err = r.redisCli.Get(channel + ":key").Result(); err == redis.Nil {
			newKey, err = r.SetKey(channel)
			log.Debugf("[KEY] new channel [%s]: %s", channel, newKey)
			return
		}
		return
	}

	var key interface{}
	var found bool
	if key, found = r.localCache.Get(channel); found {
		return key.(string), nil
	}
	newKey, err = r.SetKey(channel)
	log.Debugf("[KEY] new channel [%s]: %s", channel, newKey)
	return
}

func (r *RoomKeysType) GetChannel(key string) (channel string, err error) {
	if !saveInLocal {
		return r.redisCli.Get("key:" + key).Result()
	}

	chann, found := r.localCache.Get(key)
	if found {
		return chann.(string), nil
	} else {
		return "", fmt.Errorf("%s does not exists", key)
	}
}

func (r *RoomKeysType) DeleteChannel(channel string) bool {
	if !saveInLocal {
		return r.redisCli.Del(channel+":key").Err() != nil
	}

	key, ok := r.localCache.Get(channel)
	if ok {
		r.localCache.Delete(channel)
		r.localCache.Delete(key.(string))
		return true
	}
	return false
}

func (r *RoomKeysType) DeleteKey(key string) bool {
	if !saveInLocal {
		return r.redisCli.Del("key:"+key).Err() != nil
	}

	channel, ok := r.localCache.Get(key)
	if ok {
		r.localCache.Delete(channel.(string))
		r.localCache.Delete(key)
		return true
	}
	return false
}

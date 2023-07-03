package cache

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/labstack/gommon/log"
	"time"
)

var (
	cache *bigcache.BigCache
	err   error
)

func Init() {
	config := bigcache.Config{
		// number of shards (must be a power of 2)
		Shards: 1024,

		// time after which entry can be evicted
		LifeWindow: 15 * time.Minute,

		// Interval between removing expired entries (clean up).
		// If set to <= 0 then no action is performed.
		// Setting to < 1 second is counterproductive — bigcache has a one second resolution.
		CleanWindow: 5 * time.Minute,

		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,

		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,

		// prints information about additional memory allocation
		Verbose: true,
	}

	cache, err = bigcache.New(context.Background(), config)
	if err != nil {
		log.Fatal("bigcache init error:", err)
	}
}

func Get(key string) ([]byte, error) {
	return cache.Get(key)
}

func Set(key string, entry []byte) error {
	return cache.Set(key, entry)
}

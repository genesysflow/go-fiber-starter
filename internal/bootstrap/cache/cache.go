package cache

import (
	"github.com/genesysflow/go-fiber-starter/utils/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type Cache struct {
	Log  zerolog.Logger
	Cfg  *config.Config
	Conn *redis.Conn
}

func NewCache(cfg *config.Config, log zerolog.Logger) *Cache {
	c := &Cache{
		Cfg: cfg,
		Log: log,
	}

	return c
}

func (c *Cache) ConnectCache() {
	opt, err := redis.ParseURL(c.Cfg.Cache.Redis.URL)
	if err != nil {
		c.Log.Error().Err(err).Msg("An error occurred while trying to parse the cache url")
	}

	rdb := redis.NewClient(opt)
	c.Conn = rdb.Conn()
}

func (c *Cache) ShutdownCache() {
	err := c.Conn.Close()
	if err != nil {
		c.Log.Error().Err(err).Msg("An error occurred while trying to shutdown the cache connection")
	} else {
		c.Log.Info().Msg("Cache connection shutdown successfully")
	}
}

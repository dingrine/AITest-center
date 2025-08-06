package redis

import (
	"context"

	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	Client *redis.Client
}

const Nil = redis.Nil

func New(logger log.Logger, cfg *Config) *Client {
	cfg.OnConnect = func(ctx context.Context, cn *redis.Conn) error {
		logger.Infof("redis connected")
		return nil
	}

	client := redis.NewClient(cfg)
	return &Client{client}
}

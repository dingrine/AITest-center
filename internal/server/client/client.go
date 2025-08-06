package client

import (
	"context"
	"examCenter/internal/chaindata/models"
	"examCenter/pkg/redis"

	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/trace/jaeger"
	"gorm.io/driver/mysql"

	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
	"gorm.io/gorm"

	// "github.com/CrazyThursdayV50/pkgo/store/db/gorm"
	"github.com/CrazyThursdayV50/pkgo/trace"
)

type Clients struct {
	// DB     *gorm.DB
	DB     *gorm.DB
	Redis  *redis.Client
	Tracer trace.TracerCreator
}

func New(ctx context.Context, cfg *Config, logger log.Logger) *Clients {
	var client Clients
	{
		logger := defaultlogger.New(cfg.Jaeger.Log)
		logger.Init()
		tracer, err := jaeger.New(ctx, cfg.Jaeger.Jaeger, logger)
		if err != nil {
			panic(err)
		}

		client.Tracer = tracer
	}
	// {
	// 	logger := defaultlogger.New(cfg.Mysql.Log)
	// 	logger.Init()
	// 	db := gorm.NewDB(logger, client.Tracer.NewTracer("db"), cfg.Mysql.Gorm)
	// 	models.SetSchema(cfg.Mysql.Gorm.Schema)
	// 	client.DB = db
	// }
	{
		db, err := gorm.Open(mysql.Open(cfg.Mysql.Gorm.DSN), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		models.SetSchema(cfg.Mysql.Gorm.Schema)
		client.DB = db
	}
	{
		client.Redis = redis.New(logger, cfg.Redis)
	}
	return &client
}

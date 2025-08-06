package client

import (
	"examCenter/pkg/redis"

	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
	"github.com/CrazyThursdayV50/pkgo/store/db/gorm"
	"github.com/CrazyThursdayV50/pkgo/telegram"
	"github.com/CrazyThursdayV50/pkgo/trace/jaeger"
)

type Config struct {
	Mysql    *Mysql
	Redis    *redis.Config
	Telegram *telegram.Config
	Jaeger   *Jaeger
}

type Mysql struct {
	Gorm *gorm.Config
	Log  *defaultlogger.Config
}

type Jaeger struct {
	Log    *defaultlogger.Config
	Jaeger *jaeger.Config
}

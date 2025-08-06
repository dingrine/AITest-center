package server

import (
	"examCenter/internal/core"
	"examCenter/internal/core/model"
	"examCenter/internal/server/client"
	"examCenter/internal/server/service"

	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
)

type Config struct {
	Log     *defaultlogger.Config
	Client  *client.Config
	Service *service.Config
	Model   *model.Config
	Core    *core.Config

	// Service
}

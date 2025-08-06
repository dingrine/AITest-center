package service

import (
	"examCenter/internal/core"
	"examCenter/internal/server/repository"
	"examCenter/internal/service/http"
	"examCenter/internal/service/http/api"
	"fmt"

	"github.com/CrazyThursdayV50/pkgo/log"
)

type Services struct {
	logger log.Logger
	cfg    *Config
	Http   http.Service
}

func New(cfg *Config, logger log.Logger, repos *repository.Repositories, triggers *core.Triggers) *Services {
	fmt.Println("logger", logger)
	return &Services{
		logger: logger,
		cfg:    cfg,
		Http:   api.New(logger, repos, triggers),
	}
}

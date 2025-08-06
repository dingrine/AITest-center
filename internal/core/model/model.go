package model

import (
	"os"

	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/go-deepseek/deepseek"
)

type Model struct {
	Logger      log.Logger
	DsClient    deepseek.Client
	cfg         Config
	DeepSeekKey string
	OpenAIKey   string
}

type DetraderParams struct {
	ContentHash string `json:"content_hash"`
	Question    string `json:"question"`
	TestId      int64  `json:"test_id"`
	TraderId    int64  `json:"trader_id"`
}

func New(cfg *Config, logger log.Logger) *Model {
	var model = Model{
		Logger:      logger,
		cfg:         *cfg,
		DeepSeekKey: os.Getenv(cfg.DeepSeekKey),
		OpenAIKey:   os.Getenv(cfg.OpenAIKey),
	}
	model.DsClient = model.InitDeepseek()
	return &model
}

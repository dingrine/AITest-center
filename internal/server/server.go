package server

import (
	"context"
	"examCenter/internal/chaindata/models"
	"examCenter/internal/core"
	"examCenter/internal/core/model"
	"examCenter/internal/server/client"
	"examCenter/internal/server/repository"
	"examCenter/internal/server/service"
	"examCenter/pkg/json"
	"flag"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/CrazyThursdayV50/pkgo/config"
	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log"
	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	logger log.Logger

	clients  *client.Clients
	services *service.Services
	repos    *repository.Repositories
	core     *core.Core
	model    *model.Model
	cfg      *Config
}

var configDir string
var configName string

func init() {
	flag.StringVar(&configDir, "d", ".", "config dir")
	flag.StringVar(&configName, "c", "config", "config name, without extension")
}

func New() *Server {
	var cfg, err = config.GetConfig[Config](configDir, configName, "yml")
	if err != nil {
		panic(err)
	}
	var s Server
	// 初始化上下文
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.cfg = cfg
	return &s
}

func (s *Server) init() {
	json.Init()
	s.logger = defaultlogger.New(s.cfg.Log)
	s.logger.Init()
	// 连接数据库
	s.clients = client.New(s.ctx, s.cfg.Client, s.logger)
	s.repos = repository.New(s.clients)
	s.model = model.New(s.cfg.Model, s.logger)
	s.core = core.New(s.ctx, s.logger, s.cfg.Core, s.model, s.clients)
	// 初始化路由
	s.services = service.New(s.cfg.Service, s.logger, s.repos, s.core.Triggers())
	// 初始化参数
}

func (s *Server) Run() {
	s.init()
	// 数据库迁移
	// 添加context
	err := s.clients.DB.Migrator().AutoMigrate(
		new(models.ExamInfo),
	// new(models.QAList),
	// new(models.Certification),
	// new(models.ResultDetail),
	)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup

	s.services.Serve(s.ctx, &wg)
	s.core.Run(s.ctx)

	goo.Go(func() {
		defer s.cancel()
		var signalPipe = make(chan os.Signal, 1)
		signal.Notify(signalPipe, os.Interrupt, syscall.SIGTERM)
		<-signalPipe
	})

	wg.Wait()
}

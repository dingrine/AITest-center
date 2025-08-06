package api

import (
	"examCenter/internal/core"
	"examCenter/internal/server/repository"
	"fmt"

	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/gin-gonic/gin"
)

var getRouter = make(map[string]func(*gin.Context))
var postRouter = make(map[string]func(*gin.Context))

func register(router map[string]func(*gin.Context), path string, handler func(*gin.Context)) {
	fmt.Println("path is ", path)
	router[path] = handler
}

func (s *Service) register() {
	register(postRouter, "/exam/create", s.CreateExam)
	register(getRouter, "/exam/history/list/query/:detraderId", s.QueryExamListByExamId)
	register(getRouter, "/exam/status/query/:examId", s.QueryExamStatusByExamId)
}

func New(logger log.Logger, Repos *repository.Repositories, triggers *core.Triggers) *Service {
	var s = Service{
		logger:   logger,
		RepoData: Repos.RepoExamData,
		triggers: triggers,
	}

	s.register()
	return &s
}

func (s *Service) RegisterRouter(router *gin.RouterGroup) {
	for k, v := range getRouter {
		s.logger.Printf("kv is", k, v)
		router.GET(k, v)
	}
	for k, v := range postRouter {
		router.POST(k, v)
	}
}

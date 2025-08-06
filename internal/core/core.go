package core

import (
	"context"
	"examCenter/internal/chaindata/models"
	"examCenter/internal/core/model"
	"examCenter/internal/server/client"
	"examCenter/internal/server/repository"
	"fmt"
	"sync"

	"github.com/CrazyThursdayV50/pkgo/goo"
	"github.com/CrazyThursdayV50/pkgo/log"
)

type Core struct {
	logger      log.Logger
	ctx         context.Context
	cfg         Config
	repos       *repository.Repositories
	clients     *client.Clients
	EvaSystem   []int
	RegionType  []int
	QA          []QA
	model       *model.Model
	regionQPath string

	triggers *Triggers
}

type Triggers struct {
	newExamStart chan *models.ExamInfo
}

func (t *Triggers) NewExamStart(exam *models.ExamInfo) {
	t.newExamStart <- exam
}

func (c *Core) Triggers() *Triggers {
	return c.triggers
}

func New(context context.Context, logger log.Logger, cfg *Config, model *model.Model, clients *client.Clients) *Core {
	return &Core{
		logger:   logger,
		ctx:      context,
		cfg:      *cfg,
		model:    model,
		clients:  clients,
		triggers: &Triggers{newExamStart: make(chan *models.ExamInfo)},
	}
}

func (c *Core) Run(ctx context.Context) {
	var examNewPipe = make(chan *models.ExamInfo, 100)
	var ansNewPipe = make(chan *QA, 100)
	var reportNewPipe = make(chan *QA, 100)
	wg := &sync.WaitGroup{}
	c.logger.Printf("begin")
	goo.Go(func() {
		for exam := range c.triggers.newExamStart {
			fmt.Println("token is ", exam)
			examNewPipe <- exam
		}
	})
	// 生成题目
	c.GeneratorQuestion()
	c.StartExam(ctx, examNewPipe, ansNewPipe, wg)
	c.CalculateScore(ctx, ansNewPipe, reportNewPipe, wg)
	c.ExamReportGenerator(ctx, reportNewPipe)
	// c.ExamInProgress(ch)
}

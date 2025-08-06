package core

import (
	"context"
	"examCenter/internal/core/propmt"
	"fmt"
	"sync"

	"github.com/CrazyThursdayV50/pkgo/worker"
	"github.com/go-deepseek/deepseek/request"
)

func (c *Core) CalculateScore(ctx context.Context, ch <-chan *QA, chQa chan<- *QA, wg *sync.WaitGroup) {
	worker, _ := worker.New("calculateScore", func(exam *QA) {
		c.logger.Printf("CalculateScore running")
		userMessage := fmt.Sprintf("编号: %s; 题目: %s; DeTrader答案: %s; 标准答案: %s; 维度: %s; 题目难度: %s", exam.QCode, exam.Q, exam.DetraderAns, exam.Ans, exam.Capability, exam.Level)
		var message = []*request.Message{{
			Role:    "system",
			Content: propmt.EvaluationPropmt,
		}, {
			Role:    "user",
			Content: userMessage,
		}}
		c.logger.Printf("message is %s\n, userMessage is %s", userMessage)
		exam.Analysis = c.model.RunDeepseek(message)
		c.repos.RepoExamData.AddQARecord(ctx, exam.ExamId, exam.QCode, exam.Q, exam.DetraderAns, exam.Analysis)
		c.logger.Printf("analysis is%s:%s", exam.QCode, exam.Analysis)

		chQa <- exam
		wg.Wait()

		close(chQa)

	})
	worker.WithContext(ctx)
	worker.WithLogger(c.logger)
	worker.WithGraceful(true)
	worker.WithTrigger(ch)
	worker.Run()
}

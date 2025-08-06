package core

import (
	"context"
	"examCenter/internal/chaindata/models"

	"github.com/CrazyThursdayV50/pkgo/worker"
)

func (c *Core) ExamInProgress(ctx context.Context, ch <-chan *models.ExamInfo) {

	// 看看能不能把题目当做chan传到gorutine中，异步考试

	// 根据examId查询该考试的考试状态，如果在进行中则证明上次考试未考完中止，则查询该examId的考试问题id进行到哪一个问题
	// 循环考试问题，起始题目从考试问题id开始，如果考试问题id为空，则考试的起始题目从第一题开始。
	// 将考试问题及选项以json的格式传给agent进行考试
	// agent返回回答后，将考试id、问题类别、问题、标准答案、选项及agent回答存到问答表中，同时QA存储agent回答
	// 结束后调用分析报告函数，对agent内容进行分析

	// for _, item := range c.QA {
	// 	q := &QA{Q: item.Q, Opts: item.Opts}
	// 	qStr, err := json.Marshal(q)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	var message = []*request.Message{
	// 		{Role: "system", Content: "你是一个答题专家，请根据题目和选项进行回答"},
	// 		{Role: "user", Content: string(qStr)},
	// 	}
	// 	c.model.RunDeepseek(message)
	// }
	worker, _ := worker.New("startExam", func(exam *models.ExamInfo) {
		if exam.ExamCount > 1 {
			return
		}

	})

	worker.WithContext(ctx)
	worker.WithLogger(c.logger)
	worker.WithGraceful(true)
	worker.WithTrigger(ch)
	worker.Run()
}

package core

import (
	"context"
	"examCenter/internal/chaindata/models"
	"fmt"
	"sync"

	"github.com/CrazyThursdayV50/pkgo/worker"
)

func (c *Core) StartExam(ctx context.Context, ch <-chan *models.ExamInfo, qa chan<- *QA, wg *sync.WaitGroup) {

	// 看看能不能把题目当做chan传到gorutine中，异步考试

	// 根据examId查询该考试的考试状态，如果在进行中则证明上次考试未考完中止，则查询该examId的考试问题id进行到哪一个问题
	// 循环考试问题，起始题目从考试问题id开始，如果考试问题id为空，则考试的起始题目从第一题开始。
	// 将考试问题及选项以json的格式传给agent进行考试
	// agent返回回答后，将考试id、问题类别、问题、标准答案、选项及agent回答存到问答表中，同时QA存储agent回答
	// 结束后调用分析报告函数，对agent内容进行分析

	worker, _ := worker.New("startExam", func(exam *models.ExamInfo) {
		c.logger.Printf("StartExam begin %+v", ch)
		// if exam.CertificationStatus != models.EXAM_UNVERIFIED && exam.CertificationStatus != models.EXAM_VERIFIED_EXPIRED {
		// 	return
		// }
		c.logger.Printf("qa length is", len(c.QA))
		for i, item := range c.QA {
			wg.Add(1)
			go func(qItem QA, index int) {
				defer func() {
					wg.Done()
				}()
				q := fmt.Sprintf("问题: %s, 选项: %s", qItem.Q, qItem.Opts)

				c.logger.Printf("qa loop processing")
				// 调用模型获取answer
				answer, err := c.model.RunDetrader(q, exam.Id, exam.DetraderId)
				if err != nil {
					panic(err)
				}
				c.logger.Printf("q is %s, answer is %s", q, answer)
				// 需要qTime&aTime
				resultChan := &QA{ExamId: qItem.ExamId, QCode: qItem.QCode, Q: q, DetraderAns: answer, Ans: qItem.Ans, Capability: qItem.Capability, Level: qItem.Level}
				qa <- resultChan
			}(item, i)
		}

		// for r := range resultChan {
		// 	// 将qa存入数据库
		// 	// c.repos.RepoExamData.AddQARecord(ctx, exam.Id, r.QCode, r.Q, r.DetraderAns)
		// 	qa <- &r
		// }

	})

	worker.WithContext(ctx)
	worker.WithLogger(c.logger)
	worker.WithGraceful(true)
	worker.WithTrigger(ch)
	worker.Run()
}

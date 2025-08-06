package core

import (
	"context"
	"encoding/json"
	"examCenter/internal/core/propmt"
)

func (c *Core) ExamReportGenerator(ctx context.Context, ch <-chan *QA) {
	var analysis []string
	// 生成分析报告
	for qa := range ch {
		analysis = append(analysis, qa.Analysis)
	}

	analysisData, err := json.Marshal(analysis)
	if err != nil {
		panic(err)
	}

	result := c.model.RunGPT(string(analysisData), propmt.EvaluationPropmt)

	c.logger.Printf("report result is ", result)
	// 将分析报告的内容存到数据库（更新考试信息表）
	// 判断考试是否通过
	// 考试通过：将考试信息存到考试详情表、认证证书表
	// 考试不通过：
	// 查询考试信息表拿到考试次数
	// 1.判断本次考试是第几次考试，如果是第一次考试，则调用考试函数进行重考
	// 2.如果是第二次考试则退出，用户付费后进行重考
}

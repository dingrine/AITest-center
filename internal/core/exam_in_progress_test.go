package core

import (
	"examCenter/internal/core/model"
	"examCenter/internal/core/propmt"
	"fmt"
	"testing"

	"github.com/CrazyThursdayV50/pkgo/log"
	defaultlogger "github.com/CrazyThursdayV50/pkgo/log/default"
	"github.com/go-deepseek/deepseek/request"
)

func TestExamProgress(t *testing.T) {
	var c = Core{regionQPath: "/Users/dingding/Documents/项目/ai-test-center/assets/json",
		QA: []QA{{Q: "根据投资者适当性原则，金融机构在推荐产品时，以下哪项行为违反了职业伦理？ ( )  ", Ans: "B", Opts: "A. 基于客户风险承受能力匹配低风险产品 , B. 优先推荐高佣金产品而忽略客户需求 , C. 全面评估客户投资目标和财务状况, D. 向保守型投资者解释产品风险", QCode: "投资者保护与职业伦理-1", Capability: "投资者保护与职业伦理 ", Level: "中等"}}}
	var logger log.Logger
	logConfig := defaultlogger.Config{
		Development: true,
		Console:     true,
		Level:       "debug",
		CallerSkip:  1,
	}
	logger = defaultlogger.New(&logConfig)
	logger.Init()
	cfg := model.Config{DeepSeekKey: "DEEPSEEK_API_KEY",
		OpenAIKey:       "OPENAI_API_KEY",
		ServerUrl:       "http://127.0.0.1:7890",
		DetraderHashKey: "70025F77782F1C97B1681AAC0E9F8B9A842B20ED14AD38CA9",
	}
	c.model = model.New(&cfg, logger)
	// c.GeneratorQuestion()
	// if err := json.Unmarshal([]byte(""), &c.QA); err != nil {
	// 	panic(err)
	// }
	// for _, qItem := range c.QA {
	// 	q := fmt.Sprintf("提问：%s %s", qItem.Q, qItem.Opts)
	// 	fmt.Println("q", q)
	// 	examId := int64(100)
	// 	modelId := int64(1136)
	// 	c.model.RunDetrader(q, examId, modelId)
	// }
	// c.ExamInProgress()
	fmt.Println(len(c.QA))
	for _, item := range c.QA {
		q := fmt.Sprintf("问题: %s, 选项: %s", item.Q, item.Opts)
		// 需要qTime&aTime
		userMessage := fmt.Sprintf("编号: %s; 题目: %s; DeTrader答案: %s; 标准答案: %s; 维度: %s; 题目难度: %s", item.QCode, q, "b", item.Ans, item.Capability, item.Level)
		var message = []*request.Message{{
			Role:    "system",
			Content: propmt.EvaluationPropmt,
		}, {
			Role:    "user",
			Content: userMessage,
		}}
		fmt.Printf("userMessage is %s", userMessage)
		analysis := c.model.RunDeepseek(message)
		fmt.Println("analysis is", analysis)
	}
}

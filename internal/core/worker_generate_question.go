package core

import (
	"encoding/json"
	"examCenter/internal/core/propmt"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-deepseek/deepseek/request"
)

func (c *Core) GeneratorQuestion() {
	var qaJson string
	qaJson, err := c.clients.Redis.Client.Get(c.ctx, "exam:qa:Protest:Mat-hk").Result()

	if err != nil {
		c.logger.Printf("can not find q records in cache")
	}

	if qaJson != "" {
		if err := json.Unmarshal([]byte(qaJson), &c.QA); err != nil {
			c.logger.Printf("qaJson 解析错误")
			panic(err)
		}
		return
	}

	var qa []map[string]interface{}
	var result string
	qa = ExcelToJson()
	num := len(qa)/5 + 1
	fmt.Println("num", num)
	sysPropmt := propmt.GeneratorQuestionPropmt
	for i := 1; i < num+1; i++ {
		n := i * 5
		if n > len(qa) {
			n = len(qa)
		}
		qaData, err := json.Marshal(qa[(i-1)*5 : n])
		c.logger.Printf("qaData is ", string(qaData))
		if err != nil {
			panic(err)
		}

		userPrompt := fmt.Sprintf(
			`[file name]: %s
			[file content begin]
			%s
			[file content end]
			%s`, "json文件", qaData, sysPropmt)
		message := []*request.Message{
			{Role: "system", Content: sysPropmt},
			{Role: "user", Content: userPrompt},
		}
		resp := c.model.RunDeepseek(message)
		c.logger.Printf("deepseek response is ", resp)
		replacer := strings.NewReplacer(`"题目编号":`, `"QCode":`, `"难度等级":`, `"Level":`, `"能力维度":`, `"Capability":`, `"题干":`, `"Q":`, `"选项":`, `"Opts":`, `"标准答案":`, `"Ans":`, `"解析":`, `"Analysis":`, `"评分要点":`, `"RatingPoints":`)
		resp = replacer.Replace(resp)
		result = fmt.Sprintf("%s%s", resp, result)
	}

	re := regexp.MustCompile(`(?s)\[\s*{.*?}\s*]`)
	match := re.FindString(result)
	if match != "" {
		err := c.clients.Redis.Client.Set(c.ctx, "exam:qa:Protest:Mat-hk", match, 0).Err()
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal([]byte(match), &c.QA); err != nil {
			panic(err)
		}

	} else {
		c.logger.Fatal("no question")
	}
	c.logger.Printf("c.QA is ", c.QA)
}

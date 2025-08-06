package model

import (
	"context"
	"fmt"

	"github.com/go-deepseek/deepseek"
	"github.com/go-deepseek/deepseek/request"
)

func (m *Model) InitDeepseek() deepseek.Client {
	var (
		err    error
		client deepseek.Client
	)

	config := deepseek.NewConfigWithDefaults()
	config.ApiKey = m.DeepSeekKey
	config.TimeoutSeconds = 30000

	client, err = deepseek.NewClientWithConfig(config)
	if err != nil {
		panic(err)
	}

	return client
}

func (m *Model) RunDeepseek(message []*request.Message) string {
	f := float32(0.6)
	temperature := &f

	chetReq := &request.ChatCompletionsRequest{
		Model:       deepseek.DEEPSEEK_REASONER_MODEL,
		Temperature: temperature,
		Messages:    message,
	}

	// m.Logger.Printf("deepseek response is running")
	ctx := context.Background()
	m.Logger.Printf("deepseek response is running, and message is %+v", chetReq)
	resp, err := m.DsClient.CallChatCompletionsReasoner(ctx, chetReq)

	// m.Logger.Printf("deepseek response success")
	fmt.Println("deepseek response success")

	m.Logger.Printf("resp is %s, err is %s", resp, resp.Choices, resp.Choices[0].Message, resp.Choices[0].Message.Content)
	if err != nil {
		m.Logger.Printf("deepseek response error is %s", err)
		panic(err)
	}

	if resp == nil || len(resp.Choices) == 0 {
		m.Logger.Error("resp is null")
		return ""
	}

	return resp.Choices[0].Message.Content
}

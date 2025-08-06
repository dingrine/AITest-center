package model

import (
	"context"
	"net/http"
	"net/url"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func (m *Model) RunGPT(userPrompt, systemPropmt string) string {
	proxyURL, err := url.Parse(m.cfg.ServerUrl)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	httpClient := &http.Client{Transport: transport}
	client := openai.NewClient(
		option.WithAPIKey(m.OpenAIKey),
		option.WithHTTPClient(httpClient),
	)
	resp, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemPropmt),
			openai.UserMessage(userPrompt),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		panic(err.Error())
	}
	// println("chatCompletion.Choices[0].Message.Content = ", chatCompletion.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content

}

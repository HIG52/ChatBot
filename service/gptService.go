package service

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const key = "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"

func ChatBot(cht string) string {

	return chatServer(cht)

}

// 챗봇 연결
func chatServer(cht string) string {
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: cht,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	ans := resp.Choices[0].Message.Content

	return ans
}

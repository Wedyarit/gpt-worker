package main

import (
	"context"
	"fmt"
	"gpt-worker/internal/executor"
	"gpt-worker/internal/functions"
	"gpt-worker/internal/utils"
	"gpt-worker/pkg/api"

	"github.com/sashabaranov/go-openai"
)

func main() {
	ctx := context.Background()
	client := api.NewClient()

	dialogue := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleUser, Content: "Read current directory"},
	}

	openaiTools := utils.ConvertToOpenAITools(functions.AllFunctions)

	for {
		resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo16K,
			Messages: dialogue,
			Tools:    openaiTools,
		})

		if err != nil || len(resp.Choices) != 1 {
			fmt.Printf("Error during completion: err:%v len(choices):%v\n", err, len(resp.Choices))
			break
		}

		msg := resp.Choices[0].Message
		dialogue = append(dialogue, msg)

		if len(msg.ToolCalls) == 0 {
			break
		}

		dialogue = executor.ExecuteFunction(dialogue, msg, functions.AllFunctions)
	}

	fmt.Printf("OpenAI answered the original request with: %v\n", dialogue[len(dialogue)-1].Content)
}

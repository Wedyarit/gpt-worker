package executor

import (
	"fmt"
	"gpt-worker/internal/interfaces"

	"github.com/sashabaranov/go-openai"
)

func ExecuteFunction(dialogue []openai.ChatCompletionMessage, msg openai.ChatCompletionMessage, functions []interfaces.Function) []openai.ChatCompletionMessage {
	for _, function := range functions {
		if function.Definition().Name == msg.ToolCalls[0].Function.Name {
			fmt.Printf("Invoking %v function with arguments: %v\n", function.Definition().Name, msg.ToolCalls[0].Function.Arguments)
			result := function.Execute(msg.ToolCalls[0].Function.Arguments)
			dialogue = append(dialogue, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleTool,
				Content:    result,
				Name:       function.Definition().Name,
				ToolCallID: msg.ToolCalls[0].ID,
			})
			break
		}
	}

	return dialogue
}

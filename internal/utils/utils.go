package utils

import (
	"gpt-worker/internal/interfaces"

	"github.com/sashabaranov/go-openai"
)

func ConvertToOpenAITools(functions []interfaces.Function) []openai.Tool {
	var tools []openai.Tool
	for _, f := range functions {
		tools = append(tools, openai.Tool{
			Type:     openai.ToolTypeFunction,
			Function: *f.Definition(),
		})
	}
	return tools
}

// Admob Facebook Appsflyer Unity ads | Удалять файлы (папки) которые взыываются

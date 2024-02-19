package interfaces

import "github.com/sashabaranov/go-openai"

type Function interface {
	Execute(arguments string) string
	Definition() *openai.FunctionDefinition
}

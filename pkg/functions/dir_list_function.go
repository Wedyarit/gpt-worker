package functions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ListDirFunction struct {
	*openai.FunctionDefinition
}

func NewListDirFunction() *ListDirFunction {
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"directory": {
				Type:        jsonschema.String,
				Description: "The directory to list files and directories in",
			},
		},
		Required: []string{"directory"},
	}
	f := openai.FunctionDefinition{
		Name:        "list_directory",
		Description: "List files and directories in the given directory",
		Parameters:  params,
	}
	return &ListDirFunction{FunctionDefinition: &f}
}

func (ldf *ListDirFunction) Execute(arguments string) string {
	var params map[string]interface{}
	err := json.Unmarshal([]byte(arguments), &params)
	if err != nil {
		return fmt.Sprintf("error unmarshalling arguments: %v", err)
	}
	directory, ok := params["directory"].(string)
	if !ok {
		return "Invalid directory"
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Sprintf("error reading directory: %v", err)
	}

	var result string
	for _, file := range files {
		result += file.Name() + "\n"
	}

	return result
}

func (ldf *ListDirFunction) Definition() *openai.FunctionDefinition {
	return ldf.FunctionDefinition
}

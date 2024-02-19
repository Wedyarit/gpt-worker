package functions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ReadFileFunction struct {
	*openai.FunctionDefinition
}

func NewReadFileFunction() *ReadFileFunction {
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"file_path": {
				Type:        jsonschema.String,
				Description: "The path to the file to read",
			},
		},
		Required: []string{"file_path"},
	}
	f := openai.FunctionDefinition{
		Name:        "read_file",
		Description: "Read the content of the specified file",
		Parameters:  params,
	}
	return &ReadFileFunction{FunctionDefinition: &f}
}

func (rdf *ReadFileFunction) Execute(arguments string) string {
	var params map[string]interface{}
	err := json.Unmarshal([]byte(arguments), &params)
	if err != nil {
		return fmt.Sprintf("error unmarshalling arguments: %v", err)
	}
	filePath, ok := params["file_path"].(string)
	if !ok {
		return "Invalid file path"
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Sprintf("error reading file: %v", err)
	}

	return string(content)
}

func (rdf *ReadFileFunction) Definition() *openai.FunctionDefinition {
	return rdf.FunctionDefinition
}

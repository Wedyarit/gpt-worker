package functions

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type CreateFileFunction struct {
	*openai.FunctionDefinition
}

func NewCreateFileFunction() *CreateFileFunction {
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"file_path": {
				Type:        jsonschema.String,
				Description: "The path to the file to create",
			},
			"content": {
				Type:        jsonschema.String,
				Description: "The content to write to the file",
			},
		},
		Required: []string{"file_path", "content"},
	}
	f := openai.FunctionDefinition{
		Name:        "create_file",
		Description: "Create a file with the specified content",
		Parameters:  params,
	}
	return &CreateFileFunction{FunctionDefinition: &f}
}

func (cff *CreateFileFunction) Execute(arguments string) string {
	var params map[string]interface{}
	err := json.Unmarshal([]byte(arguments), &params)
	if err != nil {
		return fmt.Sprintf("error unmarshalling arguments: %v", err)
	}
	filePath, ok := params["file_path"].(string)
	if !ok {
		return "Invalid file path"
	}
	content, ok := params["content"].(string)
	if !ok {
		return "Invalid content"
	}

	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Sprintf("error creating file: %v", err)
	}

	return fmt.Sprintf("File %s created successfully", filePath)
}

func (cff *CreateFileFunction) Definition() *openai.FunctionDefinition {
	return cff.FunctionDefinition
}

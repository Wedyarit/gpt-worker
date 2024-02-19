package functions

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type UnzipFunction struct {
	*openai.FunctionDefinition
}

func NewUnzipFunction() *UnzipFunction {
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"zip_path": {
				Type:        jsonschema.String,
				Description: "The path to the zip file to unzip",
			},
			"output_dir": {
				Type:        jsonschema.String,
				Description: "The path to the directory to unzip into",
			},
		},
		Required: []string{"zip_path", "output_dir"},
	}
	f := openai.FunctionDefinition{
		Name:        "unzip",
		Description: "Unzip the specified zip file into the specified directory",
		Parameters:  params,
	}
	return &UnzipFunction{FunctionDefinition: &f}
}

func (uzf *UnzipFunction) Execute(arguments string) string {
	var params map[string]interface{}
	err := json.Unmarshal([]byte(arguments), &params)
	if err != nil {
		return fmt.Sprintf("error unmarshalling arguments: %v", err)
	}
	zipPath, ok := params["zip_path"].(string)
	if !ok {
		return "Invalid zip path"
	}
	outputDir, ok := params["output_dir"].(string)
	if !ok {
		return "Invalid output directory"
	}

	err = unzip(zipPath, outputDir)
	if err != nil {
		return fmt.Sprintf("error unzipping file: %v", err)
	}

	return fmt.Sprintf("Zip file %s unzipped successfully to directory %s", zipPath, outputDir)
}

func unzip(zipPath, outputDir string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		filePath := filepath.Join(outputDir, file.Name)

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, file.Mode())
			if err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (uzf *UnzipFunction) Definition() *openai.FunctionDefinition {
	return uzf.FunctionDefinition
}

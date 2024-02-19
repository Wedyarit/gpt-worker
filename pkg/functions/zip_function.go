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

type ZipFunction struct {
	*openai.FunctionDefinition
}

func NewZipFunction() *ZipFunction {
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"files": {
				Type:        jsonschema.Array,
				Description: "The paths to the files to zip",
				Items: &jsonschema.Definition{
					Type: jsonschema.String,
				},
			},
			"output_path": {
				Type:        jsonschema.String,
				Description: "The path to the output zip file",
			},
		},
		Required: []string{"files", "output_path"},
	}
	f := openai.FunctionDefinition{
		Name:        "zip_files",
		Description: "Zip the specified files into a zip archive",
		Parameters:  params,
	}
	return &ZipFunction{FunctionDefinition: &f}
}

func (zff *ZipFunction) Execute(arguments string) string {
	var params map[string]interface{}
	err := json.Unmarshal([]byte(arguments), &params)
	if err != nil {
		return fmt.Sprintf("error unmarshalling arguments: %v", err)
	}
	files, ok := params["files"].([]interface{})
	if !ok {
		return "Invalid files"
	}
	outputPath, ok := params["output_path"].(string)
	if !ok {
		return "Invalid output path"
	}

	zipFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Sprintf("error creating zip file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		filePath, ok := file.(string)
		if !ok {
			return "Invalid file path"
		}

		err := addFileToZip(zipWriter, filePath)
		if err != nil {
			return fmt.Sprintf("error adding file to zip: %v", err)
		}
	}

	return fmt.Sprintf("Files zipped successfully to %s", outputPath)
}

func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}

	header.Name = filepath.Base(filePath)

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return nil
}

func (zff *ZipFunction) Definition() *openai.FunctionDefinition {
	return zff.FunctionDefinition
}

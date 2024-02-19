
# GPT Worker

GPTWorker is a project template designed to provide a set of tools for working with ChatGPT function calling. It includes functions for creating, deleting, reading files, listing directories, zipping directories, and unzipping directories.

![Go Version](https://img.shields.io/badge/Go-1.22.0-blue) ![Version](https://img.shields.io/badge/Version-0.1.0-blue) [![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) 


## Example of usage

```go
dialogue := []openai.ChatCompletionMessage{
  {Role: openai.ChatMessageRoleUser, Content: "Read current directory"},
}
```

```bash
❯ go run cmd/main.go

Invoking list_directory function with arguments: {
  "directory": "."
}
OpenAI answered the original request with: The current directory contains the following files and directories:

- .DS_Store
- README.md
- cmd
- go.mod
- go.sum
- internal
- pkg
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/wedyarit/gpt-worker
```

Go to the project directory

```bash
  cd gpt-worker
```

Run the application

```bash
  go run cmd/main.go
```

Build the application

```bash
  go run cmd/main.go
```


**Do not forget to set OPENAI_API_KEY env variable**
```bash
  export OPENAI_API_KEY="YOUR_OPENAI_API_KEY"
```
## API Reference

#### Function Interface

You can create custom functions by implementing the Function interface:


```go
type Function interface {
  Execute(arguments string) string
  Definition() *openai.FunctionDefinition
}
```

To add your custom function to GPTWorker, you'll need to add the implementation to the functions_instances.

GPTWorker automatically adds this function to function_calling and GPT will be able to use this function when executing certain prompts.
## Authors

- [@wedyarit](https://www.github.com/wedyarit)


## License

[MIT](https://choosealicense.com/licenses/mit/)


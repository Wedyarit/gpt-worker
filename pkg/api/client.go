package api

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Client struct {
	*openai.Client
}

func NewClient() *Client {
	return &Client{
		Client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
	}
}

func (c *Client) CreateChatCompletion(ctx context.Context, req openai.ChatCompletionRequest) (response openai.ChatCompletionResponse, err error) {
	return c.Client.CreateChatCompletion(ctx, req)
}

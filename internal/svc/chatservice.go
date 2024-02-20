// Package contains services for interacting with LLMs (openai so far)
package svc

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type AIClientIface interface {
	CreateChatCompletion(context.Context, openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error)
}

type ChatServiceIface interface {
	DumpHistory() []ChatMessage
	InvokeUserQuery(query string) (*openai.ChatCompletionMessage, error)
	Prompt(prompt string)
}

type ChatService struct {
	client AIClientIface
	ctx     context.Context
	history ChatHistoryServiceIface
	model   string
}

// Creates a system message
func (c *ChatService) Prompt(prompt string) {
	c.history.AddToHistory(c.history.CreateSystemMessage(prompt))
}

func (c *ChatService) DumpHistory() []ChatMessage {
	return c.history.GetHistory()
}

func (c *ChatService) InvokeUserQuery(query string) (*openai.ChatCompletionMessage, error) {
	res, err := c.client.CreateChatCompletion(c.ctx, openai.ChatCompletionRequest{
		Model:       c.model,
		Temperature: 0,
		Messages:    c.makeChatCompletionMessages(query),
	})
	if err != nil {
		return nil, err
	}

	aiMessage := c.history.CreateAIMessage(res.Choices[0].Message.Content)
	c.history.AddToHistory(aiMessage)

	return &res.Choices[0].Message, nil
}

func (c *ChatService) makeChatCompletionMessages(query string) []openai.ChatCompletionMessage {
	msgs := make([]openai.ChatCompletionMessage, 0)
	userMessage := c.history.CreateUserMessage(query)
	c.history.AddToHistory(userMessage)
	for _, msg := range c.history.GetHistory() {
		msgs = append(msgs, openai.ChatCompletionMessage{Role: msg.Role, Content: msg.Content})
	}
	return msgs
}

func NewChatService(ctx context.Context, client AIClientIface, model string, history ChatHistoryServiceIface) ChatServiceIface {
	return &ChatService{
		client:  client,
		ctx:     ctx,
		history: history,
		model:   model,
	}
}

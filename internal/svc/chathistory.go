package svc

import "github.com/sashabaranov/go-openai"

type ChatMessage struct {
	Role    string
	Content string
}

type ChatHistoryServiceIface interface {
	CreateUserMessage(text string) ChatMessage
	CreateSystemMessage(text string) ChatMessage
	CreateAIMessage(text string) ChatMessage
	AddToHistory(msg ChatMessage)
	GetHistory() []ChatMessage
}

type ChatHistoryService struct {
	History []ChatMessage
}

func NewChatHistoryService() ChatHistoryServiceIface {
	return &ChatHistoryService{History: []ChatMessage{}}
}

func (chs ChatHistoryService) CreateUserMessage(text string) ChatMessage {
	return ChatMessage{Role: openai.ChatMessageRoleUser, Content: text}
}

func (chs ChatHistoryService) CreateSystemMessage(text string) ChatMessage {
	return ChatMessage{Role: openai.ChatMessageRoleSystem, Content: text}
}

func (chs ChatHistoryService) CreateAIMessage(text string) ChatMessage {
	return ChatMessage{Role: openai.ChatMessageRoleAssistant, Content: text}
}

func (chs *ChatHistoryService) AddToHistory(msg ChatMessage) {
	chs.History = append(chs.History, msg)
}

func (chs *ChatHistoryService) GetHistory() []ChatMessage {
	return chs.History
}

func (chs *ChatHistoryService) ClearHistory() {
	chs.History = []ChatMessage{}
}

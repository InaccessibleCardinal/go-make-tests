// test generated by openai

package chat

import (
	"reflect"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestChatHistoryService_CreateUserMessage(t *testing.T) {
	chs := NewChatHistoryService().(*ChatHistoryService)
	expectedMessage := ChatMessage{Role: openai.ChatMessageRoleUser, Content: "test"}
	userMessage := chs.CreateUserMessage("test")
	if !reflect.DeepEqual(userMessage, expectedMessage) {
		t.Errorf("CreateUserMessage did not return the expected message. Expected: %v, Got: %v", expectedMessage, userMessage)
	}
}

func TestChatHistoryService_CreateSystemMessage(t *testing.T) {
	chs := NewChatHistoryService().(*ChatHistoryService)
	expectedMessage := ChatMessage{Role: openai.ChatMessageRoleSystem, Content: "test"}
	systemMessage := chs.CreateSystemMessage("test")
	if !reflect.DeepEqual(systemMessage, expectedMessage) {
		t.Errorf("CreateSystemMessage did not return the expected message. Expected: %v, Got: %v", expectedMessage, systemMessage)
	}
}

func TestChatHistoryService_CreateAIMessage(t *testing.T) {
	chs := NewChatHistoryService().(*ChatHistoryService)
	expectedMessage := ChatMessage{Role: openai.ChatMessageRoleAssistant, Content: "test"}
	aiMessage := chs.CreateAIMessage("test")
	if !reflect.DeepEqual(aiMessage, expectedMessage) {
		t.Errorf("CreateAIMessage did not return the expected message. Expected: %v, Got: %v", expectedMessage, aiMessage)
	}
}

func TestChatHistoryService_AddToHistory(t *testing.T) {
	chs := NewChatHistoryService().(*ChatHistoryService)
	msg := ChatMessage{Role: "test", Content: "test"}
	chs.AddToHistory(msg)
	history := chs.GetHistory()
	if len(history) != 1 || !reflect.DeepEqual(history[0], msg) {
		t.Errorf("AddToHistory did not add the message to history correctly. History: %v", history)
	}
}

func TestChatHistoryService_ClearHistory(t *testing.T) {
	chs := NewChatHistoryService().(*ChatHistoryService)
	msg := ChatMessage{Role: "test", Content: "test"}
	chs.AddToHistory(msg)
	chs.ClearHistory()
	history := chs.GetHistory()
	if len(history) != 0 {
		t.Errorf("ClearHistory did not clear the history. History: %v", history)
	}
}

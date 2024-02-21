package chat

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

func GetChatService() ChatServiceIface {
	ctx := context.Background()
	client := openai.NewClient(os.Getenv("OPEN_AI_API_KEY"))
	return NewChatService(ctx, client, openai.GPT3Dot5Turbo, NewChatHistoryService())
}

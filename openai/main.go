//go:build mage

package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

// Quickstart passes a prompt to the Text completion API (usage: mage quickstart 'Write a tagline for an ice cream shop.')
func Quickstart(prompt string) error {
	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		return errors.New("OPENAI_API_KEY not set")
	}
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return err
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return nil

}

// Env displays sample environment variables (use: mage env > env.sh)
func Env() {
	txt := `export OPENAI_API_KEY=''
`
	fmt.Printf(txt)
}

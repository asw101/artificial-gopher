//go:build mage

package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/element-of-surprise/azopenai"
	"github.com/element-of-surprise/azopenai/auth"
	"github.com/element-of-surprise/azopenai/clients/chat"
)

// Completions is a sample request to the Completions API
func Completions() error {
	c, err := clientFromEnv("text-davinci-003")
	if err != nil {
		return err
	}
	client := c.Completions()
	ctx := context.Background()
	prompts := []string{"Once upon a time"}
	resp, err := client.Call(ctx, prompts)
	if err != nil {
		return err
	}
	fmt.Println(resp.Text[0])
	return nil
}

// Chat is a sample request to the Chat API
func Chat() error {
	c, err := clientFromEnv("gpt-35-turbo")
	if err != nil {
		return err
	}
	client := c.Chat()
	ctx := context.Background()
	messages := []chat.SendMsg{
		{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
		{
			Role:    "user",
			Content: "Does Azure OpenAI support customer managed keys?",
		},
	}
	resp, err := client.Call(ctx, messages)
	if err != nil {
		return err
	}
	fmt.Println(resp.Text[0])
	return nil
}

// Embeddings is a sample request to the Embeddings API
func Embeddings() error {
	c, err := clientFromEnv("text-embedding-ada-002")
	if err != nil {
		return err
	}
	client := c.Embeddings()
	ctx := context.Background()
	text := []string{"The food was delicious and the waiter..."}
	resp, err := client.Call(ctx, text)
	if err != nil {
		return err
	}
	fmt.Printf("%v", resp.Results)
	return nil
}

func clientFromEnv(deploymentID string) (*azopenai.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("API_KEY is not set")
	}
	resourceName := os.Getenv("RESOURCE_NAME")
	if resourceName == "" {
		return nil, errors.New("RESOURCE_NAME is not set")
	}
	if deploymentID == "" {
		deploymentID = os.Getenv("DEPLOYMENT_ID")
	}
	if deploymentID == "" {
		return nil, errors.New("DEPLOYMENT_ID is not set")
	}
	client, err := azopenai.New(resourceName, deploymentID, auth.Authorizer{ApiKey: apiKey})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Env displays sample environment variables (use: mage env > env.sh)
func Env() {
	txt := `export API_KEY=''
export RESOURCE_NAME=''
export DEPLOYMENT_ID=''
`
	fmt.Print(txt)
}

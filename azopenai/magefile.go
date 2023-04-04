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
	c, err := clientFromEnv()
	if err != nil {
		return err
	}
	client := c.Completions()
	ctx := context.Background()
	prompts := []string{"The capital of Tennessee is"}
	resp, err := client.Call(ctx, prompts)
	if err != nil {
		return err
	}
	fmt.Println(resp.Text[0])
	return nil
}

// Chat is a sample request to the Chat API
func Chat() error {
	c, err := clientFromEnv()
	if err != nil {
		return err
	}
	client := c.Chat()
	ctx := context.Background()
	messages := []chat.SendMsg{
		{
			Content: "Is Go better than Python for the use case of distributed systems?",
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
	c, err := clientFromEnv()
	if err != nil {
		return err
	}
	client := c.Embeddings()
	ctx := context.Background()
	text := []string{"Go is the best language"}
	resp, err := client.Call(ctx, text)
	if err != nil {
		return err
	}
	fmt.Printf("%v", resp.Results)
	return nil
}

func clientFromEnv() (*azopenai.Client, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("API_KEY is not set")
	}
	resourceName := os.Getenv("RESOURCE_NAME")
	if resourceName == "" {
		return nil, errors.New("RESOURCE_NAME is not set")
	}
	deploymentName := os.Getenv("DEPLOYMENT_NAME")
	if deploymentName == "" {
		return nil, errors.New("DEPLOYMENT_NAME is not set")
	}
	client, err := azopenai.New(resourceName, deploymentName, auth.Authorizer{ApiKey: apiKey})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Env displays sample environment variables (use: mage env > env.sh)
func Env() {
	txt := `export API_KEY=''
export RESOURCE_NAME=''
export DEPLOYMENT_NAME=''
`
	fmt.Printf(txt)
}

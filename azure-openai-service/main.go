//go:build mage

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Quickstart takes a prompt parameter and runs the Azure OpenAI Service quickstart
func Quickstart(prompt string) error {
	b, err := quickstart(prompt)
	var tmp any
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	b, err = json.MarshalIndent(tmp, "", "    ")
	if err != nil {
		return err
	}
	//b, err = io.ReadAll(res.Body)
	//if err != nil {
	//	return err
	//}
	fmt.Printf("%s\n", b)
	return nil
}

// quickstart calls the Azure OpenAI Service REST endpoint per https://learn.microsoft.com/en-us/azure/cognitive-services/openai/quickstart?pivots=rest-api
func quickstart(prompt string) ([]byte, error) {
	baseURL := os.Getenv("ENDPOINT")
	if baseURL == "" {
		return nil, errors.New("ENDPOINT not set")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, errors.New("API_KEY not set")
	}
	deploymentName := os.Getenv("DEPLOYMENT_NAME")
	if deploymentName == "" {
		return nil, errors.New("DEPLOYMENT_NAME not set")
	}

	url1 := fmt.Sprintf("%s/openai/deployments/%s/completions?api-version=2022-12-01", baseURL, deploymentName)
	dic1 := map[string]string{
		"prompt": prompt,
	}
	b, err := json.Marshal(dic1)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url1, bytes.NewBuffer(b))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", apiKey)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Env displays sample environment variables (use: mage env > env.sh)
func Env() {
	txt := `export ENDPOINT=''
export API_KEY=''
export DEPLOYMENT_NAME=''
`
	fmt.Printf(txt)
}

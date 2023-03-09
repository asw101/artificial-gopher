# artificial-gopher

### Requirements

- [Go](https://go.dev/dl/)
- [Mage](https://magefile.org/) (`go install github.com/magefile/mage@latest`)

Note: these snippets do not _require_ Mage, but we use this as a runner for our project. Copy the relevant code to your own project to run without Mage.

## azure-openai-service

[Docs (learn.microsoft.com)](https://learn.microsoft.com/en-us/azure/cognitive-services/openai/quickstart?pivots=rest-api)

```bash
$ mage
Targets:
  env           displays sample environment variables (use: mage env > env.sh)
  quickstart    takes a prompt parameter and runs the Azure OpenAI Service quickstart
```

### Usage

```bash
cd azure-openai-service/
mage env > env.sh
# edit env.sh and update ENDPOINT, API_KEY, DEPLOYMENT_NAME
mage quickstart 'Once upon a time'
```

## openai

[Docs (platform.openai.com)](https://platform.openai.com/docs/api-reference/introduction) | [sashabaranov/openai-go (github.com)](https://github.com/sashabaranov/go-openai#go-openai)

```bash
$ mage
Targets:
  env           displays sample environment variables (use: mage env > env.sh)
  quickstart    passes a prompt to the Text completion API (usage: mage quickstart 'Write a tagline for an ice cream shop.')
```


### Usage

```bash
cd openai/
mage env > env.sh
# edit env.sh and update OPENAI_API_KEY
mage quickstart 'Write a tagline for an ice cream shop.'
```
# azopenai

```
$ mage
Targets:
  chat           is a sample request to the Chat API
  completions    is a sample request to the Completions API
  embeddings     is a sample request to the Embeddings API
  env            displays sample environment variables (use: mage env > env.sh)
```

## Usage

```bash
cd azopenai/
mage env > env.sh
# edit env.sh and update API_KEY, RESOURCE_NAME, DEPLOYMENT_NAME
source env.sh

mage completions

mage chat

mage embeddings
```

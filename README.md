# grafana-webhook-to-telegram

Accepts a Grafana webhook and forwards the text to Telegram.

## ENV variables

Environment variables are you can set

```dotenv
HTTP_SERVER_LISTEN_ADDR=127.0.0.1:8080
TELEGRAM_API_HOST=https://api.telegram.org
# BOT_API_KEY_<NAME> — suffix matches /api/<bot_name>/... in the URL (uppercase); hyphens in the name become underscores (my-bot → BOT_API_KEY_MY_BOT)
BOT_API_KEY_MYBOT=1234567890:***********************************
```

## Usage

1. Grafana **Contact point** (HTTP) URL:

   `POST` or `PUT` to `http://<host>:<port>/api/<bot_name>/<chat_id>`

   `chat_id` is the chat or channel ID (e.g. from [@userinfobot](https://t.me/userinfobot) or the Bot API).

2. Request body is Grafana webhook JSON (`message`, `title`, `status`). Telegram receives `message`, or `title` if `message` is empty.

## Docker build and run

Build the image from the repository root:

```bash
docker build -t grafana-webhook-to-telegram .
```

Run with environment variables:

```bash
docker run --rm -p 8080:8080 \
  -e BOT_API_KEY_MYBOT='your_token' \
  grafana-webhook-to-telegram
```

Override the listen address if needed:

```bash
docker run --rm -p 8080:8080 \
  -e HTTP_SERVER_LISTEN_ADDR=0.0.0.0:8080 \
  -e BOT_API_KEY_MYBOT='your_token' \
  grafana-webhook-to-telegram
```

Locally without Docker: `go run ./cmd` (`.env` is not loaded automatically — export variables or use something like `env $(grep -v '^#' .env | xargs) go run ./cmd`).

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this does

Receives Grafana webhook POSTs/PUTs at `POST|PUT /api/{bot_name}/{chat_id}` and forwards the `message` (or `title` if message is empty) to a Telegram chat via a bot token looked up from environment.

## Commands

```bash
# Run tests
go test ./...

# Vet
go vet ./...

# Run single test
go test ./internal/handler/ -run TestWebhook

# Run locally (dotenv is auto-loaded from .env if the file exists)
go run ./cmd

# Run locally with env vars explicitly
env $(grep -v '^#' .env | xargs) go run ./cmd

# Docker build & run
docker build -t grafana-webhook-to-telegram .
docker run --rm -p 8080:8080 -e BOT_API_KEY_MYBOT='token' grafana-webhook-to-telegram
```

## Architecture

```
cmd/main.go              — entry point: loads .env, wires dependencies, starts HTTP server
internal/config/         — env-based config via caarlos0/env; LOG_ prefix for log settings
internal/storage/        — APIKeyStorage interface; APIKeyENVStorage resolves BOT_API_KEY_<NAME> env vars
internal/handler/        — HTTP handler; parses Grafana JSON, calls MessageSender interface
internal/telegram/       — Telegram Bot API client implementing MessageSender
```

**Key interface**: `handler.MessageSender` (in `internal/handler/webhook.go`) decouples the handler from Telegram — use this for mocking in tests.

**Bot name → env var mapping**: URL segment `my-bot` → env var `BOT_API_KEY_MY_BOT` (hyphens become underscores, uppercased). Implemented in `storage/storage.go`.

## Environment variables

| Variable | Default | Notes |
|---|---|---|
| `BOT_API_KEY_<NAME>` | — | Required per bot; NAME is uppercased URL segment |
| `HTTP_SERVER_LISTEN_ADDR` | `127.0.0.1:8080` | Docker image defaults to `0.0.0.0:8080` |
| `TELEGRAM_API_HOST` | `https://api.telegram.org` | Override for proxy/test |
| `HTTPS_PROXY` / `HTTP_PROXY` | — | Standard Go proxy env vars |
| `LOG_MODE` | `development` | `development` or `production` (zap) |
| `LOG_LEVEL` | `info` | zap level |
| `LOG_DISABLE_STACKTRACE` | `true` | |

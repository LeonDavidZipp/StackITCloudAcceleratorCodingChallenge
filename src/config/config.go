package config

import "os"

var (
	TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")
	TELEGRAM_CHAT_ID   = os.Getenv("TELEGRAM_CHAT_ID")
)

const (
	HTTP_SERVER_PORT string = "8080"
	GRPC_SERVER_PORT string = "8088"
)

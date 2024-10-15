package main

import "github.com/saintbyte/telegram_gateway_api"

func main() {
	tga := telegram_gateway_api.NewTelegramGatewayAPI()
	tga.SendVerificationMessage()
}

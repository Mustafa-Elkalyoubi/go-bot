package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/mustafa-elkalyoubi/discord-bot/bot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot.Run()
}

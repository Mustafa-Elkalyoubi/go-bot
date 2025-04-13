package events

import (
	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(sesh *discordgo.Session, e *discordgo.Ready) {
	Logger.Println("Bot Connected")
}

package events

import (
	"github.com/bwmarrin/discordgo"
)

func MessageCreateHandler(sesh *discordgo.Session, msg *discordgo.MessageCreate) {
	Logger.Println(msg.Content)
}

package events

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var Logger *log.Logger

func RegisterHandlers(sesh *discordgo.Session) {
	Logger = log.New(os.Stdout, "bot(events)", log.LstdFlags)

	sesh.AddHandlerOnce(ReadyHandler)
	sesh.AddHandler(MessageCreateHandler)
}

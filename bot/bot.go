package bot

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"

	"github.com/mustafa-elkalyoubi/discord-bot/bot/events"
)

var Logger = log.New(os.Stdout, "bot", log.LstdFlags)

func Run() {
	token := os.Getenv("DISCORD_TOKEN")
	sesh, err := discordgo.New("Bot " + token)

	if err != nil {
		log.Fatal(err)
	}

	events.RegisterHandlers(sesh)
	sesh.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)
	discordgo.IntentsAllWithoutPrivileged

	openErr := sesh.Open()

	if openErr != nil {
		Logger.Fatalf("Cannot open session: %v", openErr)
	}

	defer sesh.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	Logger.Println("Shutting down...")
}

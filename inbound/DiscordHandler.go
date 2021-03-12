package inbound

import (
	"Blibots/service"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var messageParserService *service.MessageParserService
var session *discordgo.Session

func InitDiscordHandler() *discordgo.Session {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

	if err != nil {
		panic("Error starting discord bot: " + err.Error())
	}

	messageParserService = service.MessageParserServiceObj
	session = discord
	return session
}

func Listen() {
	session.AddHandler(messageCreate)

	err := session.Open()

	if err != nil {
		log.Fatal("error", err.Error())
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	(*messageParserService).ParseMessage(s ,m.Content, m.ChannelID)
}

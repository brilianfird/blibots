package service

import "github.com/bwmarrin/discordgo"

var MessageParserServiceObj = new(MessageParserService)

type MessageParserService interface {
	ParseMessage(session *discordgo.Session, msg string, channelId string)
}

func initializeMessageParserService() {
	*MessageParserServiceObj = newMessageParserServiceImpl()
}

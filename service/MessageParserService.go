package service

import "github.com/bwmarrin/discordgo"

var MessageParserServiceObj *MessageParserService

type MessageParserService interface {
	ParseMessage(session *discordgo.Session, msg string, channelId string)
}

func initializeMessageParserService() {
	var messageParserService MessageParserService = newMessageParserServiceImpl()
	MessageParserServiceObj = &messageParserService
}
package outbound

import "github.com/bwmarrin/discordgo"

var DiscordOutboundObj *DiscordOutbound

type DiscordOutbound interface {
	SendMessage(session *discordgo.Session, message string, channelId string) error
}

func initializeDiscordOutbound() {
	var discordOutbound DiscordOutbound = newDiscordOutboundImpl()
	DiscordOutboundObj = &discordOutbound
}

package outbound

import "github.com/bwmarrin/discordgo"

type DiscordOutboundImpl struct {

}

func newDiscordOutboundImpl() *DiscordOutboundImpl {
	return &DiscordOutboundImpl{}
}

func (d DiscordOutboundImpl) SendMessage(session *discordgo.Session, message string, channelId string) error {
	_, err := session.ChannelMessageSend(channelId, message)

	return err
}
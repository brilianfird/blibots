package inbound

func InitializeInbound() {
	InitDiscordHandler()
	InitializeCronJob()
	Listen()
}

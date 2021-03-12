package main

import (
	"Blibots/inbound"
	"Blibots/outbound"
	"Blibots/repository"
	"Blibots/service"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	repository.InitRepository()
	outbound.InitializeOutbound()
	service.InitializeService()
	inbound.InitializeInbound()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}



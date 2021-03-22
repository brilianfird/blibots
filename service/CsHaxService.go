package service

import (
	"Blibots/model"
	"github.com/bwmarrin/discordgo"
)

var CsHaxServiceObj *CsHaxService = new(CsHaxService)

type CsHaxService interface {
	Insert(csHax model.CsHax) (*model.CsHax, error)
	GetAll() ([]model.CsHax, error)
	CheckBans(session *discordgo.Session)
}

func initCsHaxService() {
	*CsHaxServiceObj = newCsHaxServiceImpl()
}

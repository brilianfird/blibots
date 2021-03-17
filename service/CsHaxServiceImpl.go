package service

import (
	"Blibots/model"
	"Blibots/outbound"
	"Blibots/repository"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly/v2"
	"strings"
)

type CsHaxServiceImpl struct {
	csHaxRepository *repository.CsHaxRepository
	discordOutbound *outbound.DiscordOutbound
}

func newCsHaxServiceImpl() *CsHaxServiceImpl {
	return &CsHaxServiceImpl{csHaxRepository: repository.CsHaxRepositoryObj,
		discordOutbound: outbound.DiscordOutboundObj}
}

func (c CsHaxServiceImpl) Insert(csHax model.CsHax) (*model.CsHax, error) {
	return (*c.csHaxRepository).Insert(csHax)
}

func (c CsHaxServiceImpl) GetAll() ([]model.CsHax, error) {
	return (*c.csHaxRepository).GetAll()
}

func (c CsHaxServiceImpl) CheckBans(session *discordgo.Session) {
	bannedFalse, _ := (*c.csHaxRepository).GetAllWithBannedFalse()

	for _, v := range bannedFalse {
		checkVACBan(v.SteamUrl, c, session)
	}
}

func checkVACBan(steamUrl string, csHaxService CsHaxServiceImpl, session *discordgo.Session) {
	c := colly.NewCollector()

	c.OnHTML("html", func(element *colly.HTMLElement) {
		text := element.Text
		if strings.Contains(text, "VAC ban on record") {
			(*csHaxService.csHaxRepository).UpdateBannedToTrue(steamUrl)
			(*csHaxService.discordOutbound).SendMessage(session, steamUrl+" is banned", "819965414368870401")
		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(steamUrl)
}

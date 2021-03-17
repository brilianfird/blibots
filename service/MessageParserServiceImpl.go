package service

import (
	"Blibots/model"
	"Blibots/outbound"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
	"strings"
	"time"
)

type MessageParserServiceImpl struct {
	reminderService *ReminderService
	discordOutbound *outbound.DiscordOutbound
	csHaxService    *CsHaxService
}

func newMessageParserServiceImpl() *MessageParserServiceImpl {
	return &MessageParserServiceImpl{
		csHaxService:    CsHaxServiceObj,
		reminderService: ReminderServiceObj,
		discordOutbound: outbound.DiscordOutboundObj,
	}
}

func (m MessageParserServiceImpl) ParseMessage(session *discordgo.Session, msg string, channelId string) {
	s := strings.Split(msg, " ")

	// $bliboys[0] remind[1] every[2] 1[3] day[4] for[5] 30[6] days[7] "message"[8]
	if s[0] == "$bliboys" {
		if s[1] == "remind" {
			parseRemind(m, session, s, channelId)
		} else if s[1] == "csHax" {
			//blibots cxHax add {url}
			parseCsHax(m, session, s, channelId)
		} else {
			(*m.discordOutbound).SendMessage(session, "Example: $bliboys remind every 1 day for 30 days message", channelId)
		}
	}
}

func parseCsHax(m MessageParserServiceImpl, session *discordgo.Session, s []string, channelId string) {
	//blibots cxHax add {url}
	if s[2] == "add" {
		csHax := model.CsHax{
			SteamUrl: s[3],
			Banned:   false,
		}
		_, err := (*m.csHaxService).Insert(csHax)

		if err != nil {
			log.Println("Error inserting csHax", err)
			(*m.discordOutbound).SendMessage(session, "Failed inserting csHax", channelId)
		} else {
			(*m.discordOutbound).SendMessage(session, "Success inserting csHax", channelId)
		}
	} else if s[2] == "ps" {
		all, _ := (*m.csHaxService).GetAll()

		for _, v := range all {
			(*m.discordOutbound).SendMessage(session, "SteamUrl: "+v.SteamUrl+" banned: "+strconv.FormatBool(v.Banned), channelId)
		}
	}
}

func parseRemind(m MessageParserServiceImpl, session *discordgo.Session, s []string, channelId string) {
	interval, err := strconv.Atoi(s[3])

	if err != nil {
		fmt.Println("error", err)
		(*m.discordOutbound).SendMessage(session, "Fail", channelId)
	}

	howLong, err := strconv.Atoi(s[6])

	if err != nil {
		fmt.Println("error", err)
		(*m.discordOutbound).SendMessage(session, "Fail", channelId)
	}
	howMuch := howLong / interval

	combinedMessage := ""
	for i := 8; i < len(s); i++ {
		combinedMessage = combinedMessage + s[i] + " "
	}
	var reminderList []model.Reminder
	for i := 1; i < howMuch; i++ {
		remindTime := time.Now().Local().Add(time.Hour*time.Duration(24*i) +
			time.Minute +
			time.Second)
		reminder := model.Reminder{
			Remind:     combinedMessage,
			RemindTime: remindTime,
			ChannelId:  channelId,
		}
		reminderList = append(reminderList, reminder)
	}

	(*m.reminderService).InsertMany(reminderList)
	(*m.discordOutbound).SendMessage(session, "Noted Boys", channelId)
}

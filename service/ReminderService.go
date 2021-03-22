package service

import (
	"Blibots/model"
	"github.com/bwmarrin/discordgo"
)

var ReminderServiceObj = new(ReminderService)

type ReminderService interface {
	InsertMany(reminders []model.Reminder) ([]model.Reminder, error)
	ProcessAndDeleteReminderWithRemindTimeLessThanNow(session *discordgo.Session)
}

func initializeReminderService() {
	*ReminderServiceObj = NewReminderServiceImpl()
}

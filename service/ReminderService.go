package service

import (
	"Blibots/model"
	"github.com/bwmarrin/discordgo"
)

var ReminderServiceObj *ReminderService

type ReminderService interface {
	InsertMany(reminders []model.Reminder) ([]model.Reminder, error)
	ProcessAndDeleteReminderWithRemindTimeLessThanNow(session *discordgo.Session)
}

func initializeReminderService() {
	var reminderServiceImpl ReminderService = NewReminderServiceImpl()
	ReminderServiceObj = &reminderServiceImpl
}



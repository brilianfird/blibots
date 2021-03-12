package service

import (
	"Blibots/model"
	"Blibots/outbound"
	"Blibots/repository"
	"github.com/bwmarrin/discordgo"
)

type ReminderServiceImpl struct {
	reminderRepository *repository.ReminderRepository
	discordOutbound *outbound.DiscordOutbound
}

func NewReminderServiceImpl() *ReminderServiceImpl {
	return &ReminderServiceImpl{
		reminderRepository: repository.ReminderRepositoryObj,
		discordOutbound: outbound.DiscordOutboundObj,
	}
}

func (r ReminderServiceImpl) InsertMany(reminders []model.Reminder) ([]model.Reminder, error) {
	reminderRepository := *(r.reminderRepository)
	return reminderRepository.InsertMany(reminders)
}

func (r ReminderServiceImpl) ProcessAndDeleteReminderWithRemindTimeLessThanNow(session *discordgo.Session) {
	reminderRepository := *(r.reminderRepository)
	reminders := reminderRepository.FindWithRemindTimeLessThanNow()

	for i := 0;i < len(reminders);i++ {
		(*r.discordOutbound).SendMessage(session, reminders[0].Remind, reminders[0].ChannelId)
	}

	reminderRepository.DeleteWithRemindTimeLessThanNow()

}

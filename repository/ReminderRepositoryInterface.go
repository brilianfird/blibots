package repository

import (
	"Blibots/model"
)

var ReminderRepositoryObj = new(ReminderRepository)

type ReminderRepository interface {
	InsertOne(r model.Reminder) (*model.Reminder, error)
	InsertMany(r []model.Reminder) ([]model.Reminder, error)
	FindWithRemindTimeLessThanNow() []model.Reminder
	DeleteWithRemindTimeLessThanNow()
	Close()
}

func InitializeReminderRepository() {
	*ReminderRepositoryObj = ReminderRepository(getSQLLite())
}

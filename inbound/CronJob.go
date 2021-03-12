package inbound

import (
	"Blibots/service"
	"github.com/robfig/cron"
	"log"
	"time"
)

var reminderService *service.ReminderService

func InitializeCronJob() {
	reminderService = service.ReminderServiceObj
	c := cron.New()
	c.AddFunc("0 * * * * *", checkScheduler)

	c.Start()
}

func checkScheduler() {
	log.Println("scheduler is running at", time.Now())
	(*reminderService).ProcessAndDeleteReminderWithRemindTimeLessThanNow(session)

}

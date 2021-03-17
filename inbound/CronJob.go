package inbound

import (
	"Blibots/service"
	"github.com/robfig/cron"
	"log"
	"time"
)

var reminderService *service.ReminderService
var csHaxService *service.CsHaxService

func InitializeCronJob() {
	reminderService = service.ReminderServiceObj
	csHaxService = service.CsHaxServiceObj
	c := cron.New()
	c.AddFunc("0 * * * * *", checkScheduler)
	c.AddFunc("0 0 * * * *", checkCsHax)
	c.Start()
}

func checkScheduler() {
	log.Println("reminder scheduler is running at", time.Now())
	(*reminderService).ProcessAndDeleteReminderWithRemindTimeLessThanNow(session)
}

func checkCsHax() {
	log.Println("csHax scheduler is running at", time.Now())
	(*csHaxService).CheckBans(session)
}

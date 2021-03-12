package model

import "time"

type Reminder struct {
	Remind string
	RemindTime time.Time
	ChannelId string
}

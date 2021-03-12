package repository

import (
	"Blibots/model"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

type SqlLiteRepository struct {
	SqlLiteDB *sql.DB
}

func getSQLLite() *SqlLiteRepository {
	workDirectory, _ := os.Getwd()
	sqlLiteDB, err := sql.Open("sqlite3", workDirectory+"\\db\\reminder.db")
	if err != nil {
		sqlLiteDB.Close()
		log.Fatal(err.Error())
	}

	return &SqlLiteRepository{
		SqlLiteDB: sqlLiteDB,
	}
}

func (repository *SqlLiteRepository) InsertOne(reminder model.Reminder) (*model.Reminder, error) {

	insertReminderSQL := `INSERT INTO reminder(remind, remind_time, channelId) VALUES (?, ?, ?)`

	prepare, err := repository.SqlLiteDB.Prepare(insertReminderSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = prepare.Exec(reminder.Remind, reminder.RemindTime.UTC().Format(time.RFC3339), reminder.ChannelId)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &reminder, nil
}

func (repository *SqlLiteRepository) InsertMany(r []model.Reminder) ([]model.Reminder, error) {
	if len(r) == 0 {
		return []model.Reminder{}, nil
	}

	for _, v  := range r {
		repository.InsertOne(v)
	}

	return r, nil
}

func (repository *SqlLiteRepository) FindWithRemindTimeLessThanNow() []model.Reminder {
	sqlQuery := `SELECT * FROM REMINDER WHERE remind_time <= ?`
	cursor, err := repository.SqlLiteDB.Query(sqlQuery, time.Now().UTC().Format(time.RFC3339))

	if err != nil {
		log.Fatal(err.Error())
	}

	var ret []model.Reminder

	for cursor.Next() {
		var id string
		var remind string
		var remindTime time.Time
		var channelId string
		err := cursor.Scan(&id, &remind, &remindTime, &channelId)

		if err != nil {
			log.Fatal(err.Error())
		}

		ret = append(ret, model.Reminder{
			Remind:     remind,
			RemindTime: remindTime,
			ChannelId:  channelId,
		})
	}

		return ret
}

func (repository *SqlLiteRepository) DeleteWithRemindTimeLessThanNow() {
	sqlQuery := `DELETE FROM reminder WHERE remind_time <= ?`

	_, err := repository.SqlLiteDB.Exec(sqlQuery, time.Now().UTC().Format(time.RFC3339))

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (repository *SqlLiteRepository) Close() {
	repository.Close()
}
package repository

import (
	"Blibots/model"
	"database/sql"
	"log"
	"os"
)

type CsHaxRepositorySQLLite struct {
	SqlLiteDB *sql.DB
}

func initCsHaxSqlLite() *CsHaxRepositorySQLLite {
	workDirectory, _ := os.Getwd()
	sqlLiteDB, err := sql.Open("sqlite3", workDirectory+"\\db\\reminder.db")
	if err != nil {
		sqlLiteDB.Close()
		log.Fatal(err.Error())
	}

	return &CsHaxRepositorySQLLite{
		SqlLiteDB: sqlLiteDB,
	}
}

func (repository CsHaxRepositorySQLLite) GetAll() ([]model.CsHax, error) {
	sqlQuery := `SELECT * FROM cs_hax`
	query, err := repository.SqlLiteDB.Query(sqlQuery)

	if err != nil {
		return []model.CsHax{}, err
	}
	var ret []model.CsHax
	for query.Next() {
		var id string
		var steamUrl string
		var banned bool
		err := query.Scan(&id, &steamUrl, &banned)

		if err == nil {
			ret = append(ret, model.CsHax{
				SteamUrl: steamUrl,
				Banned:   banned,
			})
		}
	}
	return ret, nil
}

func (repository CsHaxRepositorySQLLite) Insert(csHax model.CsHax) (*model.CsHax, error) {
	insertReminderSQL := `INSERT INTO cs_hax (steam_url, banned) VALUES (?, ?)`

	prepare, err := repository.SqlLiteDB.Prepare(insertReminderSQL)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = prepare.Exec(csHax.SteamUrl, csHax.Banned)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &csHax, nil
}

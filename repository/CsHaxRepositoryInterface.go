package repository

import "Blibots/model"

var CsHaxRepositoryObj *CsHaxRepository

type CsHaxRepository interface {
	GetAll() ([]model.CsHax, error)
	Insert(csHax model.CsHax) (*model.CsHax, error)
	GetAllWithBannedFalse() ([]model.CsHax, error)
	UpdateBannedToTrue(steamUrl string)
}

func InitCsHaxRepository() {
	var csHaxSqlLite CsHaxRepository = initCsHaxSqlLite()
	CsHaxRepositoryObj = &csHaxSqlLite
}

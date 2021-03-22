package repository

import (
	"Blibots/model"
)

var CsHaxRepositoryObj = new(CsHaxRepository)

type CsHaxRepository interface {
	GetAll() ([]model.CsHax, error)
	Insert(csHax model.CsHax) (*model.CsHax, error)
	GetAllWithBannedFalse() ([]model.CsHax, error)
	UpdateBannedToTrue(steamUrl string)
}

func InitCsHaxRepository() {
	*CsHaxRepositoryObj = CsHaxRepository(initCsHaxSqlLite())
}

package service

import "Blibots/model"

var CsHaxServiceObj *CsHaxService

type CsHaxService interface {
	Insert(csHax model.CsHax) (*model.CsHax, error)
	GetAll() ([]model.CsHax, error)
}

func initCsHaxService() {
	var csHaxService CsHaxService = newCsHaxServiceImpl()
	CsHaxServiceObj = &csHaxService
}

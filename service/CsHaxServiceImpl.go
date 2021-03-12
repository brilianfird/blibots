package service

import (
	"Blibots/model"
	"Blibots/repository"
)

type CsHaxServiceImpl struct {
	csHaxRepository *repository.CsHaxRepository
}

func newCsHaxServiceImpl() *CsHaxServiceImpl {
	return &CsHaxServiceImpl{csHaxRepository: repository.CsHaxRepositoryObj}
}

func (c CsHaxServiceImpl) Insert(csHax model.CsHax) (*model.CsHax, error) {
	return (*c.csHaxRepository).Insert(csHax)
}

func (c CsHaxServiceImpl) GetAll() ([]model.CsHax, error) {
	return (*c.csHaxRepository).GetAll()
}

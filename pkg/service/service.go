package service

import (
	"balance"
	"balance/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Billing interface {
	AddMoney(account balance.User) error
	Reserve(ord balance.Order) error
	WriteOff(ord balance.Order) error
	GetBalance(id int) (int, error)
	Dereserve(ord balance.Order) error
}

type Info interface {
	MakeReport(year, month int) error
	GiveName(serv balance.Report) error
	GetHistory(id int, sort string, p *balance.Pagination) ([]balance.History, error)
}

type Service struct {
	Billing
	Info
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Billing: NewBillingService(r.Billing),
		Info:    NewInfoService(r.Info),
	}
}

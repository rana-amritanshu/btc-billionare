package service

import (
	"btc/pkg/repository"
	"time"
)

const datetimeFormat = "2006-01-02T15:04:05+0000"

type Add struct {
	Amount   float32
	Datetime string
}

type AddRepository interface {
	Save(add *repository.Add) error
}

type AddService struct {
	repository AddRepository
}

func (a *AddService) Save(add *Add) error {
	saveData := &repository.Add{
		Amount:   add.Amount,
		Datetime: time.Now().UTC().Format(datetimeFormat),
	}

	if add.Datetime != "" {
		saveData.Datetime = add.Datetime
	}

	err := a.repository.Save(saveData)
	if err != nil {
		return err
	}
	return nil
}

func NewAddService(repo AddRepository) *AddService {
	return &AddService{repo}
}

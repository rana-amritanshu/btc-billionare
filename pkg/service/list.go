package service

import (
	"btc/pkg/repository"
)

type ListRepository interface {
	List(params *repository.WalletSearchParams) ([]*repository.Wallet, error)
}

type Wallet struct {
	Amount   float64
	Datetime string
}

type ListServiceParams struct {
	StartDatetime string
	EndDatetime   string
}

type ListService struct {
	repository ListRepository
}

func (a *ListService) List(params *ListServiceParams) ([]*Wallet, error) {

	wallets, err := a.repository.List(
		&repository.WalletSearchParams{
			StartDatetime: params.StartDatetime,
			EndDatetime:   params.EndDatetime,
		})

	if err != nil {
		return nil, err
	}

	var results []*Wallet

	for _, wallet := range wallets {
		results = append(results, &Wallet{
			Amount:   wallet.Amount,
			Datetime: wallet.Datetime,
		})
	}

	return results, nil
}

func NewListService(repo ListRepository) *ListService {
	return &ListService{repo}
}

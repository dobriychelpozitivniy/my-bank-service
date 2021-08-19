package service

import (
	"errors"
	"github.com/dobriychelpozitivniy/mybank/pkg/repository"
	"math"
)

const SBP2RUB = 0.7523
const RUB2SBP = 1.3333
const percentageOfSavings = 0.06
const withdrawalPercentage = 0.7

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) AddFunds(sum float64) error {
	cash, err := s.repo.GetCash()
	if err != nil {
		return err
	}

	result := cash + sum

	err = s.repo.UpdateCash(Round(result))
	if err != nil {
		return err
	}

	err = s.sumProfit()
	if err != nil {
		return err
	}

	return nil
}

func (s *AccountService) sumProfit() error {
	cash, err := s.repo.GetCash()
	if err != nil {
		return err
	}

	result := cash + cash * percentageOfSavings

	return s.repo.UpdateCash(Round(result))
}

func (s *AccountService) GetCurrency() (string, error) {
	return s.repo.GetCurrency()
}

func (s *AccountService) GetBalance(cur string) (float64, error) {
	cash, err := s.repo.GetCash()
	if err != nil {
		return 0, err
	}

	currencyRate, err := s.GetAccountCurrencyRate(cur)
	if err != nil {
		return 0, err
	}

	result := cash * currencyRate

	return Round(result), nil
}

func (s *AccountService) GetAccountCurrencyRate(cur string) (float64, error) {
	curAccount, err := s.GetCurrency()
	if err != nil {
		return 0, err
	}

	if (cur == curAccount) {
		return 1, nil
	} else if (cur == "SBP" && curAccount == "RUB") {
		return RUB2SBP, nil
	} else if (cur == "RUB" && curAccount == "SBP") {
		return SBP2RUB, nil
	}

	return 0, err
}

func (s *AccountService) Withdraw(sum float64) error {
	balance, err := s.repo.GetCash()
	if err != nil {
		return err
	}

	if sum > Round(balance * withdrawalPercentage) {
		return errors.New("you cannot withdraw more than 70% of the current balance")
	}

	result := balance - sum

	err = s.repo.UpdateCash(Round(result))
	if err != nil {
		return err
	}

	return nil
}

func Round(sum float64) float64 {
	return math.Round(sum*100)/100
}

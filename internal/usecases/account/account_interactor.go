package account

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/model"
	"github.com/dhiemaz/AccountService/internal/repositories/account"
)

type AccountInteractor struct {
	accountRepository account.Repository
	out               Output
}

// NewAccountInteractor object creation
func NewAccountInteractor(accountRepository account.Repository, out Output) *AccountInteractor {
	return &AccountInteractor{
		accountRepository: accountRepository,
		out:               out,
	}
}

// CreateNewAccount use case
func (i *AccountInteractor) CreateNewAccount(data model.Account) (model.Account, error) {
	result, err := i.accountRepository.InsertAccount(data)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "CreateNewAccount"}).
			Errorf("failed create a new account, error : %v", err)
	}

	return i.out.CreateNewAccountOutput(result, err)
}

// UpdateAccount use case
func (i *AccountInteractor) UpdateAccount(data model.Account) (model.Account, error) {
	result, err := i.accountRepository.UpdateAccount(data)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "UpdateAccount"}).
			Errorf("failed update account, error : %v", err)
	}

	return i.out.UpdateAccountOutput(result, err)
}

// GetAllData use case
func (i *AccountInteractor) GetAllData() ([]model.Account, error) {
	result, err := i.accountRepository.GetAllAccounts()
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllAccounts"}).
			Errorf("failed get all accounts data, error : %v", err)
	}

	return i.out.GetAllData(result, err)
}

// FindAccountByUsername use case
func (i *AccountInteractor) FindAccountByUsername(username string) (model.Account, error) {
	result, err := i.accountRepository.FindOneAccountByUsername(username)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "FindAccountByUsername"}).
			Errorf("failed find account by username, error : %v", err)
	}

	return i.out.FindAccountByUsername(result, err)
}

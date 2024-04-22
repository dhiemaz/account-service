package usecase

import (
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	"github.com/dhiemaz/AccountService/internal/domain/account/model"
	"github.com/dhiemaz/AccountService/internal/domain/account/repository"
	paginate "github.com/gobeam/mongo-go-pagination"
)

type AccountUsecase struct {
	accountRepository repository.Repository
	out               Output
}

// NewAccountUsecase object creation
func NewAccountUsecase(accountRepository repository.Repository, out Output) AccountUsecase {
	return AccountUsecase{
		accountRepository: accountRepository,
		out:               out,
	}
}

// CreateNewAccount use case
func (i AccountUsecase) CreateNewAccount(data model.Account) (model.Account, error) {
	result, err := i.accountRepository.InsertAccount(data)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "CreateNewAccount"}).
			Errorf("failed create a new account, error : %v", err)
	}

	return i.out.CreateNewAccountOutput(result, err)
}

// UpdateAccount use case
func (i AccountUsecase) UpdateAccount(data model.Account) (model.Account, error) {
	result, err := i.accountRepository.UpdateAccount(data)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "UpdateAccount"}).
			Errorf("failed update account, error : %v", err)
	}

	return i.out.UpdateAccountOutput(result, err)
}

// GetAllData use case
func (i AccountUsecase) GetAllData(page, limit int) ([]*model.Account, paginate.PaginationData, error) {
	result, pagination, err := i.accountRepository.GetAllAccounts(page, limit)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "GetAllAccounts"}).
			Errorf("failed get all accounts data, error : %v", err)
	}

	return i.out.GetAllData(result, pagination, err)
}

// FindAccountByUsername use case
func (i AccountUsecase) FindAccountByUsername(username string) (model.Account, error) {
	result, err := i.accountRepository.FindOneAccountByUsername(username)
	if err != nil {
		logger.WithFields(logger.Fields{"component": "usecase", "action": "FindAccountByUsername"}).
			Errorf("failed find account by username, error : %v", err)
	}

	return i.out.FindAccountByUsername(result, err)
}

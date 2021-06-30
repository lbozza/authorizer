package authorizer

import (
	"authorizer/entity"
	"authorizer/entity/violations"
	"authorizer/usecase/rules"
	"errors"
)

//go:generate go run github.com/golang/mock/mockgen -package=authorizer -self_package=authorizer -destination=./manager_mock.go . Manager
type Manager interface {
	CreateAccount(account entity.Account) (*entity.Account, error)
	ProcessTransaction(transaction entity.Transaction) (*entity.Account, error)
	CreateDenyList(denyList []string) (*entity.Account, error)
}

type OperationManager struct {
	initialAccount *entity.Account
	trx            []entity.Transaction
}

func NewOperationManager() *OperationManager {
	return &OperationManager{
		trx: []entity.Transaction{},
	}
}

func (m *OperationManager) CreateAccount(account entity.Account) (*entity.Account, error) {

	if m.initialAccount != nil {
		return m.initialAccount, violations.ErrAccountAlreadyInitialized
	}

	m.initialAccount = &account

	return m.initialAccount, nil
}

func (m *OperationManager) ProcessTransaction(transaction entity.Transaction) (*entity.Account, error) {
	if !m.validateAccountState() {
		return nil, errors.New("account-not-initialized")
	}

	if !m.initialAccount.ActiveCard {
		return m.initialAccount, violations.ErrCardNotActive
	}

	if !m.validateLimit(transaction) {
		return m.initialAccount, violations.ErrInsufficientLiit
	}

	rules := rules.NewRules()

	err := rules.Authorize(&transaction, m.trx, m.initialAccount.DenyList)
	if err != nil {
		return m.initialAccount, err
	}
	m.trx = append(m.trx, transaction)

	m.initialAccount.AvaliableLimit -= transaction.Amount
	return m.initialAccount, nil
}

func (m *OperationManager) CreateDenyList(denyList []string) (*entity.Account, error) {
	if !m.validateAccountState() {
		return m.initialAccount, errors.New("account-not-initialized")
	}

	m.initialAccount.DenyList = denyList

	return m.initialAccount, nil
}

func (m *OperationManager) validateLimit(transaction entity.Transaction) bool {
	account := m.initialAccount

	if account.AvaliableLimit < transaction.Amount {
		return false
	}

	return true
}

func (m *OperationManager) validateAccountState() bool{
	if m.initialAccount == nil {
		return false
	}

	return true
}
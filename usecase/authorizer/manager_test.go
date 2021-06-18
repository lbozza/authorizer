package authorizer

import (
	"authorizer/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type ManagerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl        *gomock.Controller
	handlerMock *MockManager
}

func TestManagerSuite(t *testing.T) {
	suite.Run(t, new(ManagerSuite))
}

func (s *ManagerSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.handlerMock = NewMockManager(s.ctrl)
}

func (s *ManagerSuite) ShutDown() {
	s.ctrl.Finish()
}

func (s *ManagerSuite) TestOperationManagerCreateAccount() {
	account := entity.Account{
		ActiveCard:     true,
		AvaliableLimit: 100,
	}

	operationManager := NewOperationManager()
	output, err := operationManager.CreateAccount(account)

	s.NoError(err)

	s.Equal(&account, output)
}

func (s *ManagerSuite) TestOperationManagerCreateAccountAlreadyExistsError() {
	account := entity.Account{
		ActiveCard:     true,
		AvaliableLimit: 100,
	}

	operationManager := NewOperationManager()
	operationManager.CreateAccount(account)
	_, err := operationManager.CreateAccount(account)

	s.Error(err, "account-already-initialized")
}

func (s *ManagerSuite) TestOperationManagerProcessTransaction() {
	account := entity.Account{
		ActiveCard:     true,
		AvaliableLimit: 100,
	}

	transaction := entity.Transaction{
		Merchant: "Burger King",
		Amount:   50,
		Time:     time.Now(),
	}

	expected := account
	expected.AvaliableLimit = 50

	operationManager := NewOperationManager()
	operationManager.CreateAccount(account)
	output, err := operationManager.ProcessTransaction(transaction)

	s.NoError(err)

	s.Equal(&expected, output)

}

func (s *ManagerSuite) TestOperationManagerCardNotActiveError() {
	account := entity.Account{
		ActiveCard:     false,
		AvaliableLimit: 100,
	}

	transaction := entity.Transaction{
		Merchant: "Burger King",
		Amount:   50,
		Time:     time.Now(),
	}

	operationManager := NewOperationManager()
	operationManager.CreateAccount(account)
	_, err := operationManager.ProcessTransaction(transaction)

	s.Error(err, "card-not-active")
}

func (s *ManagerSuite) TestOperationManagerInsufficientLimitError() {
	account := entity.Account{
		ActiveCard:     true,
		AvaliableLimit: 100,
	}

	transaction := entity.Transaction{
		Merchant: "Burger King",
		Amount:   110,
		Time:     time.Now(),
	}

	operationManager := NewOperationManager()
	operationManager.CreateAccount(account)
	_, err := operationManager.ProcessTransaction(transaction)

	s.Error(err, "insufficient-limit")
}

func (s *ManagerSuite) TestOperationManagerDoubleTransactionError() {
	account := entity.Account{
		ActiveCard:     true,
		AvaliableLimit: 100,
	}

	transaction := entity.Transaction{
		Merchant: "Burger King",
		Amount:   50,
		Time:     time.Now(),
	}

	operationManager := NewOperationManager()
	operationManager.CreateAccount(account)
	operationManager.ProcessTransaction(transaction)
	_, err := operationManager.ProcessTransaction(transaction)

	s.Error(err, "double-transaction")
}

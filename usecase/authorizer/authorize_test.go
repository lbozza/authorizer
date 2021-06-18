package authorizer

import (
	"authorizer/entity"
	"authorizer/processor"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AuthorizeSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller

	managerMock *MockManager
}

func TestProcessorSuite(t *testing.T) {
	suite.Run(t, new(AuthorizeSuite))
}

func (s *AuthorizeSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.managerMock = NewMockManager(s.ctrl)
}

func (s *AuthorizeSuite) ShutDown() {
	s.ctrl.Finish()
}

func (s *AuthorizeSuite) TestAuthorizeAccount() {
	input := processor.Input{
		Account: &entity.Account{
			AvaliableLimit: 1000,
			ActiveCard:     true,
		},
	}
	expected := processor.Output{
		Account: &entity.Account{
			AvaliableLimit: 1000,
			ActiveCard:     true,
		},
		Violations: []string{},
	}
	AuthorizeHandler := NewAuthorizeHandler()
	output, err := AuthorizeHandler.Handle(input)

	s.NoError(err)

	s.Equal(expected, output)
}

func (s *AuthorizeSuite) TestProcessTransaction() {
	input := processor.Input{
		Transaction: &entity.Transaction{
			Merchant: "Burger King",
			Amount:   50,
			Time:     time.Now(),
		},
	}

	inputAccount := processor.Input{
		Account: &entity.Account{
			AvaliableLimit: 1000,
			ActiveCard:     true,
		},
	}

	expected := processor.Output{
		Account: &entity.Account{
			ActiveCard:     true,
			AvaliableLimit: 950,
		},
		Violations: []string{},
	}

	authorizeHandler := NewAuthorizeHandler()
	_, _ = authorizeHandler.Handle(inputAccount)
	output, err := authorizeHandler.Handle(input)

	s.NoError(err)
	s.Equal(expected, output)
}

func (s *AuthorizeSuite) TestInvalidOperationType() {
	authorizeHandler := NewAuthorizeHandler()
	_, err := authorizeHandler.Handle(processor.Input{})

	s.Error(err, "invalid operation type")
}

func (s *AuthorizeSuite) TestAccountNotInitializedError() {
	input := processor.Input{
		Transaction: &entity.Transaction{
			Merchant: "Burger King",
			Amount:   50,
			Time:     time.Now(),
		},
	}
	expected := processor.Output{
		Account:    nil,
		Violations: []string{"account-not-initialized"},
	}

	authorizeHandler := NewAuthorizeHandler()
	output, err := authorizeHandler.Handle(input)

	s.NoError(err)
	s.Equal(expected, output)
}

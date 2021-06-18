package rules

import (
	"authorizer/entity"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var timeTransaction = time.Date(2021, 07, 11, 10, 00, 00, 00, time.UTC)

type RulesSuite struct {
	suite.Suite
	*require.Assertions

	ctrl *gomock.Controller
}

func TestProcessorSuite(t *testing.T) {
	suite.Run(t, new(RulesSuite))
}

func (s *RulesSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())
}

func (s *RulesSuite) ShutDown() {
	s.ctrl.Finish()
}

func (s *RulesSuite) TestAuthorize() {
	transaction := &entity.Transaction{
		Merchant: "Burger King",
		Amount:   110,
		Time:     time.Now(),
	}

	rules := NewRules()
	err := rules.Authorize(transaction, []entity.Transaction{})

	s.NoError(err)
}

func (s *RulesSuite) TestDoubleTransactionError() {
	transaction := &entity.Transaction{
		Merchant: "Burger King",
		Amount:   110,
		Time:     timeTransaction,
	}

	rules := NewRules()
	err := rules.Authorize(transaction, []entity.Transaction{
		{
			Merchant: "Burger King",
			Amount:   10,
			Time:     timeTransaction,
		},
	})

	s.Error(err, "double-transaction")
}

func (s *RulesSuite) TestHighFrequencySmallIntervalTransactionError() {
	transaction := &entity.Transaction{
		Merchant: "Burger King",
		Amount:   110,
		Time:     timeTransaction,
	}

	rules := NewRules()
	err := rules.Authorize(transaction, []entity.Transaction{
		{
			Merchant: "Mc Donald's",
			Amount:   10,
			Time:     timeTransaction,
		},
		{
			Merchant: "Pizza Hut",
			Amount:   10,
			Time:     timeTransaction,
		},
		{
			Merchant: "Domino's",
			Amount:   10,
			Time:     timeTransaction,
		},
	})

	s.Error(err, "high-frequency-small-interval")
}

package rules

import (
	"authorizer/entity"
	"errors"
	"strings"
	"time"
)

type DoubleTransactionRule struct {
}

const (
	timeLimit = 2 * time.Minute
)

func NewDoubleTransactionRule() *DoubleTransactionRule {
	return &DoubleTransactionRule{}
}

func (d *DoubleTransactionRule) Authorize(trx []entity.Transaction, transaction *entity.Transaction) error {
	for _, o := range trx {
		if strings.Compare(o.Merchant, transaction.Merchant) == 0 {
			duration := transaction.Time.Sub(o.Time)
			if duration < timeLimit {
				return errors.New("double-transaction")
			}
		}
	}

	return nil
}

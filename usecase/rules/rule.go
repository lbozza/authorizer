package rules

import (
	"authorizer/entity"
)

type Rules struct {
	rules *List
}

func (r *Rules) Authorize(transaction *entity.Transaction, trx []entity.Transaction) error {
	err := r.rules.DoubleTransactionRule.Authorize(trx, transaction)

	if err != nil {
		return err
	}

	err = r.rules.HighFrequencySmallIntervalRule.Authorize(trx, transaction)

	if err != nil {
		return err
	}

	return nil

}

func NewRules() *Rules {
	return &Rules{
		rules: NewRulesList(),
	}
}

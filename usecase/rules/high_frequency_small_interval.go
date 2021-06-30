package rules

import (
	"authorizer/entity"
	"authorizer/entity/violations"
)

type HighFrequencySmallIntervalRule struct {
}

func NewHighFrequencySmallIntervalRule() *HighFrequencySmallIntervalRule {
	return &HighFrequencySmallIntervalRule{}
}

func (d *HighFrequencySmallIntervalRule) Authorize(trx []entity.Transaction, transaction *entity.Transaction) error {

	sizeList := len(trx)

	if len(trx) >= 3 {
		timeDiff := transaction.Time.Sub(trx[sizeList-3].Time)

		if timeDiff < timeLimit {
			return violations.ErrHighFrequencySmallInterval
		}
	}

	return nil
}

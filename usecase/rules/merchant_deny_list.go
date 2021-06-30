package rules

import (
	"authorizer/entity"
	"errors"
	"strings"
)

type MerchantDenyList struct {

}

func NewMerchantDenyList() *MerchantDenyList {
	return &MerchantDenyList{}
}

func (m *MerchantDenyList) Authorize(merchantDenyList []string, transaction *entity.Transaction) error {

	for _, o := range merchantDenyList {
		if strings.Compare(o, transaction.Merchant) == 0 {
			return errors.New("merchant-denied")
		}
	}
	return nil
}
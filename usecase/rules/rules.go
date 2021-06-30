package rules

type List struct {
	*DoubleTransactionRule
	*HighFrequencySmallIntervalRule
	*MerchantDenyList
}

func NewRulesList() *List {
	return &List{
		NewDoubleTransactionRule(),
		NewHighFrequencySmallIntervalRule(),
		NewMerchantDenyList(),
	}
}

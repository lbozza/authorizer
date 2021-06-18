package rules

type List struct {
	*DoubleTransactionRule
	*HighFrequencySmallIntervalRule
}

func NewRulesList() *List {
	return &List{
		NewDoubleTransactionRule(),
		NewHighFrequencySmallIntervalRule(),
	}
}

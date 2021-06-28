package violations

import "errors"

var (
	ErrAccountAlreadyInitialized = errors.New("account-already-initialized")
	ErrCardNotActive = errors.New("card-not-active")
    ErrInsufficientLiit = errors.New("insufficient-limit")
    ErrDoubleTransaction = errors.New("double-transaction")
    ErrHighFrequencySmallInterval = errors.New("high-frequency-small-interval")
)
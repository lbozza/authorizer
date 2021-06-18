package authorizer

import (
	"authorizer/entity"
	"authorizer/processor"
	"errors"
)

type AuthorizeHandler struct {
	Manager
}

const (
	typeAccount     = 1
	typeTransaction = 2
)

func NewAuthorizeHandler() *AuthorizeHandler {
	return &AuthorizeHandler{NewOperationManager()}
}

func (a *AuthorizeHandler) Handle(input processor.Input) (output processor.Output, err error) {

	operationType := getOperationType(input)
	var account *entity.Account

	switch operationType {
	case typeAccount:
		account, err = a.CreateAccount(*input.Account)
	case typeTransaction:
		account, err = a.ProcessTransaction(*input.Transaction)
	case 0:
		return processor.Output{}, errors.New("invalid operation type")
	}

	violations := []string{}
	if err != nil {
		violations = append(violations, err.Error())

	}

	return processor.Output{Account: account, Violations: violations}, nil
}

func getOperationType(operation processor.Input) int {
	isAccount := operation.Account != nil
	isTransaction := operation.Transaction != nil

	if isAccount == false && isTransaction == false {
		return 0
	}
	if isAccount {
		return typeAccount
	}
	return typeTransaction
}

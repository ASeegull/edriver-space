package payment

import (
	"errors"
	"github.com/ASeegull/edriver-space/model"
)

func DoPayment(inputFunds, requiredFunds int) error {
	if inputFunds > requiredFunds {
		err := errors.New("banking service error")
		return err
	}
	if inputFunds < requiredFunds {
		err := model.ErrInsufficientFunds
		return err
	}
	return nil
}

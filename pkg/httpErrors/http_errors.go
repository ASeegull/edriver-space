package httpErrors

import (
	"fmt"
	"regexp"
	"strconv"
)

type RestErr interface {
	Status() int
	Error() string
}

type RestError struct {
	ErrStatus int    `json:"status,omitempty"`
	ErrError  string `json:"error,omitempty"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s", e.ErrStatus, e.ErrError)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func NewRestError(status int, err string) RestErr {
	return RestError{ErrStatus: status, ErrError: err}
}

func ParseRestError(e error) (RestError, error) {
	reStatus := regexp.MustCompile(`status: \d+`)
	reError := regexp.MustCompile(`errors: .+`)

	intStatus, err := strconv.Atoi(reStatus.FindString(e.Error())[8:])
	if err != nil {
		return RestError{}, err
	}
	return RestError{
		ErrStatus: intStatus,
		ErrError:  reError.FindString(e.Error()),
	}, nil
}

func ErrorResponse(err error) (int, interface{}) {
	restErr, _ := ParseRestError(err)
	return restErr.ErrStatus, restErr.ErrError
}

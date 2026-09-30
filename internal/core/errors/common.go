package errors

import "fmt"

type ErrorApp struct {
	Err     error
	Message string
}

func (e *ErrorApp) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *ErrorApp) Unwrap() error {
	return e.Err
}

func TopUpBalanceMockError() error {
	return &ErrorApp{
		Message: "Failed to top up balance (mock simulation)",
	}
}

package models

import "fmt"

type AppError struct {
	Code      int
	DescFa    string
	Private   bool
	ForceShow bool
}

func (m *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", m.Code, m.DescFa)
}

func registerAppErr(err AppError) AppError {

	return err
}

var (
	ErrUserIdMustBePostive = registerAppErr(AppError{
		Code:      1,
		DescFa:    "useename must be postive",
		ForceShow: true,
		Private:   false,
	})

	ErrPostIsInvaid = registerAppErr(AppError{
		Code:      2,
		DescFa:    "ErrPostIsInvaid",
		ForceShow: false,
		Private:   false,
	})
)

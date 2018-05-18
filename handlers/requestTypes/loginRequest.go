package requestTypes

import validator "gopkg.in/go-playground/validator.v9"

type (
	CustomValidator struct {
		Validator *validator.Validate
	}

	LoginRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	ChangePasswordRequest struct {
		UserName string `json:"username"`
	}

	ConfirmChangePasswordRequest struct {
		UserName string `json:"username"`
		COde     string `json:"code"`
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

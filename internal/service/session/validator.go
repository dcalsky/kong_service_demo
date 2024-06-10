package session

import (
	"fmt"
	"net/mail"

	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

type validator struct {
}

func (s validator) IsLegalEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s validator) ValidateLoginRequest(req dto.LoginRequest) {
	if req.Email == "" || req.Password == "" {
		panic(base.InvalidParamErr.WithRawError(fmt.Errorf("email or password cannot be empty")))
	}
}

func (s validator) ValidateRegisterRequest(req dto.RegisterRequest) {
	if req.Email == "" || req.NickName == "" || req.Password == "" {
		panic(base.InvalidParamErr.WithRawError(fmt.Errorf("email or password cannot be empty")))
	}
	if len(req.Password) < 8 || len(req.Password) > 20 {
		panic(base.PasswordLengthErr)
	}
	if !s.IsLegalEmail(req.Email) {
		panic(base.EmailFormatErr.WithRawError(fmt.Errorf("email format is invalid, email: %s", req.Email)))
	}
}

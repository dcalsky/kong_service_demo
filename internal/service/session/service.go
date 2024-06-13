package session

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/dcalsky/kong_service_demo/internal/adapter/repo"
	"github.com/dcalsky/kong_service_demo/internal/base"
	"github.com/dcalsky/kong_service_demo/internal/model/dto"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

var (
	jwtClaimsPool = sync.Pool{
		New: func() any {
			return &dto.JwtClaims{}
		},
	}
	expiresDuration = time.Hour * 7 // expires in 7 days
)

type ISessionService interface {
	generateJwt(account entity.Account) string
	Login(ctx context.Context, req dto.LoginRequest) dto.LoginResponse
	Register(ctx context.Context, req dto.RegisterRequest) dto.RegisterResponse
}

type sessionService struct {
	accountRepo repo.IAccountRepo
	secretBytes []byte
	validator
}

func NewSessionService(accountRepo repo.IAccountRepo, secret string) ISessionService {
	s := &sessionService{
		accountRepo: accountRepo,
		secretBytes: []byte(secret),
	}
	return s
}

func (s *sessionService) generateJwt(account entity.Account) string {
	claims := jwtClaimsPool.Get().(*dto.JwtClaims)
	claims.ExpiredAt = time.Now().Add(expiresDuration)
	claims.AccountId = uint(account.ID)
	claims.Email = account.Email
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.secretBytes)
	jwtClaimsPool.Put(claims)
	if err != nil {
		panic(base.InternalError.WithRawError(fmt.Errorf("sign jwt failed, err: %w", err)))
	}
	return token
}

func (s *sessionService) Login(ctx context.Context, req dto.LoginRequest) dto.LoginResponse {
	s.validator.ValidateLoginRequest(req)
	account, err := s.accountRepo.DescribeAccountByEmail(ctx, req.Email)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))
	base.PanicIf(account == nil, base.Unauthorized.WithRawError(fmt.Errorf("account not found, email: %s", req.Email)))

	if !checkPasswordHash(req.Password, account.Password) {
		panic(base.Unauthorized.WithRawError(fmt.Errorf("password not match, email: %s, account id: %d", req.Email, account.ID)))
	}

	token := s.generateJwt(*account)
	return dto.LoginResponse{
		Token:   token,
		Account: account.ToForDetail(),
	}
}

func (s *sessionService) Register(ctx context.Context, req dto.RegisterRequest) dto.RegisterResponse {
	s.validator.ValidateRegisterRequest(req)
	hashedPassword, err := hashPassword(req.Password)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	existedAccount, err := s.accountRepo.DescribeAccountByEmail(ctx, req.Email)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	if existedAccount != nil {
		panic(base.AccountEmailHasBeenTaken.WithRawError(fmt.Errorf("email: %s", req.Email)))
	}

	account := entity.NewAccount(req.Email, req.NickName, hashedPassword)
	err = s.accountRepo.CreateAccount(ctx, &account)
	base.PanicIfErr(err, base.InternalError.WithRawError(err))

	token := s.generateJwt(account)
	return dto.RegisterResponse{
		Token:   token,
		Account: account.ToForDetail(),
	}
}

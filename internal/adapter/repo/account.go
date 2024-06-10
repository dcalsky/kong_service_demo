package repo

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

type IAccountRepo interface {
	CreateAccount(ctx context.Context, account *entity.Account) error
	DescribeAccountByEmail(ctx context.Context, email string) (*entity.Account, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) IAccountRepo {
	s := &accountRepo{
		db: db,
	}
	return s
}

func (s *accountRepo) CreateAccount(ctx context.Context, account *entity.Account) error {
	return s.db.WithContext(ctx).Create(account).Error
}

func (s *accountRepo) DescribeAccountByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account entity.Account
	err := s.db.WithContext(ctx).Where("email = ?", email).First(&account).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

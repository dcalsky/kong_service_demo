package entity

import (
	"time"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

type AccountId uint

type Account struct {
	Id        AccountId `gorm:"primaryKey"`
	Email     string
	NickName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s Account) TableName() string {
	return "account"
}

func NewAccount(email, nickname, password string) Account {
	now := time.Now()
	return Account{
		Id:        0,
		Email:     email,
		NickName:  nickname,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s Account) ToForDetail() dto.AccountForDetail {
	return dto.AccountForDetail{
		Id:        uint(s.Id),
		Email:     s.Email,
		NickName:  s.NickName,
		CreatedAt: s.CreatedAt.Unix(),
		UpdatedAt: s.UpdatedAt.Unix(),
	}
}

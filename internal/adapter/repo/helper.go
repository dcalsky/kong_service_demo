package repo

import (
	"context"

	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/base"
)

// IRepoHelper wraps a transaction for service layer and extracts the transaction from context
//
//go:generate mockgen -package repo --build_flags=--mod=mod  --destination helper_mock.go . IRepoHelper
type IRepoHelper interface {
	MustTransaction(ctx context.Context, cb func(ctx context.Context))
	WithContext(ctx context.Context) *gorm.DB
}

type repoHelper struct {
	db *gorm.DB
}

func NewRepoHelper(db *gorm.DB) IRepoHelper {
	s := &repoHelper{
		db: db,
	}
	return s
}

func (s *repoHelper) WithContext(ctx context.Context) *gorm.DB {
	originTx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		return s.db.WithContext(ctx)
	}
	return originTx
}

func (s *repoHelper) MustTransaction(ctx context.Context, cb func(ctx context.Context)) {
	originTx, ok := ctx.Value(txKey).(*gorm.DB)
	if !ok {
		originTx = s.db.WithContext(ctx)
	}
	err := originTx.Transaction(func(tx *gorm.DB) error {
		exception := base.Catch(func() {
			cb(context.WithValue(ctx, txKey, tx))
		})
		// don't return exception directly due to bi-types
		if exception != nil {
			return *exception
		} else {
			return nil
		}
	})
	if err != nil {
		panic(err)
	}
}

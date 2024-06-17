package repo

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dcalsky/kong_service_demo/internal/base"
)

func TestRepoHelper_MustTransaction(t *testing.T) {
	t.Run("no exception", func(t *testing.T) {
		mockDB, mock, _ := sqlmock.New()
		mock.ExpectQuery("").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.45.1"))
		mock.ExpectBegin()
		mock.ExpectCommit()
		db, err := gorm.Open(sqlite.Dialector{Conn: mockDB}, &gorm.Config{})
		require.NoError(t, err)

		helper := NewRepoHelper(db)
		ctx := context.Background()
		helper.MustTransaction(ctx, func(ctx context.Context) {
			tx := ctx.Value(txKey)
			require.NotNil(t, tx)
			_, ok := tx.(*gorm.DB)
			require.True(t, ok)
		})
	})

	t.Run("exception", func(t *testing.T) {
		mockDB, mock, _ := sqlmock.New()
		mock.ExpectQuery("").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.45.1"))
		mock.ExpectBegin()
		mock.ExpectCommit()
		db, err := gorm.Open(sqlite.Dialector{Conn: mockDB}, &gorm.Config{})
		require.NoError(t, err)

		helper := NewRepoHelper(db)
		ctx := context.Background()
		except := base.Catch(func() {
			helper.MustTransaction(ctx, func(ctx context.Context) {
				tx := ctx.Value(txKey)
				require.NotNil(t, tx)
				_, ok := tx.(*gorm.DB)
				require.True(t, ok)

				panic(base.InternalError.WithRawError(fmt.Errorf("expected exception")))
			})
		})
		require.NotNil(t, except)
		require.Equal(t, except.Message, base.InternalError.Message)
	})
}

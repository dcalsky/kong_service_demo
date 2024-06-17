package repo

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestOrderBy(t *testing.T) {
	db := &gorm.DB{
		Statement: &gorm.Statement{
			Clauses: make(map[string]clause.Clause),
		},
	}
	BuildGormSortBy(db, []string{"-UpdatedAt", "CreatedAt"})
	clauses, ok := db.Statement.Clauses["ORDER BY"]
	require.True(t, ok)
	require.NotNil(t, clauses)
}

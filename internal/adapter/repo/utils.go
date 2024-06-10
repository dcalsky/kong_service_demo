package repo

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

const (
	sortByDescPrefix = "-"
)

func BuildGormPagination(tx *gorm.DB, option *dto.PagingOption) {
	if option == nil {
		return
	}
	tx.Offset(int(option.Skip()))
	tx.Limit(int(option.PageSize))
}

func BuildGormSortBy(tx *gorm.DB, sortBy []string) {
	if len(sortBy) == 0 {
		return
	}
	for _, sortValue := range sortBy {
		afterSortValue, found := strings.CutPrefix(sortValue, sortByDescPrefix)
		if found {
			tx.Order(clause.OrderByColumn{Column: clause.Column{Name: afterSortValue}, Desc: true})
		} else {
			tx.Order(clause.OrderByColumn{Column: clause.Column{Name: afterSortValue}, Desc: false})
		}
	}
}

package repo

import (
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/dcalsky/kong_service_demo/internal/model/dto"
)

const (
	sortByDescPrefix = "-"
	defaultPageSize  = 10
	defaultPageNum   = 1
)

func BuildGormPagination(tx *gorm.DB, option *dto.PagingOption) {
	if option == nil {
		option = &dto.PagingOption{
			PageSize: defaultPageSize,
			PageNum:  defaultPageNum,
		}
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

func BuildPageResult(tx *gorm.DB, option *dto.PagingOption) dto.PagingResult {
	if option == nil {
		option = &dto.PagingOption{
			PageSize: defaultPageSize,
			PageNum:  defaultPageNum,
		}
	}
	res := dto.PagingResult{
		PageSize: option.PageSize,
		PageNum:  option.PageNum,
		Total:    0,
	}
	var total int64
	tx.Count(&total)
	res.Total = total
	return res
}

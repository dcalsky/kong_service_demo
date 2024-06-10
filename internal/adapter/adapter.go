package adapter

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/dcalsky/kong_service_demo/internal/adapter/repo"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

var (
	db              *gorm.DB
	KongServiceRepo repo.IKongServiceRepo
	AccountRepo     repo.IAccountRepo
)

func MustInit() {
	var err error
	db, err = gorm.Open(sqlite.Open("kong.db"), &gorm.Config{
		Logger: &repo.GormLogger{},
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	KongServiceRepo = repo.NewKongServiceRepo(db.Table(entity.KongService{}.TableName()))
	AccountRepo = repo.NewAccountRepo(db.Table(entity.Account{}.TableName()))
}

func Release() {
	sqlDb, _ := db.DB()
	_ = sqlDb.Close()
}

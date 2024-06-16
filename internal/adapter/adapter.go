package adapter

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/dcalsky/kong_service_demo/internal/adapter/repo"
	"github.com/dcalsky/kong_service_demo/internal/model/entity"
)

var (
	db               *gorm.DB
	KongServiceRepo  repo.IKongServiceRepo
	AccountRepo      repo.IAccountRepo
	OrganizationRepo repo.IOrganizationRepo
	RepoHelper       repo.IRepoHelper
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
	RepoHelper = repo.NewRepoHelper(db)
	KongServiceRepo = repo.NewKongServiceRepo(RepoHelper)
	AccountRepo = repo.NewAccountRepo(RepoHelper)
	OrganizationRepo = repo.NewOrganizationRepo(RepoHelper)

	err = db.AutoMigrate(
		&entity.KongServiceVersion{},
		&entity.KongService{},
		&entity.Account{},
		&entity.Organization{},
		&entity.OrganizationAccountMapping{},
	)
	if err != nil {
		panic(fmt.Errorf("failed to migrate database: %w", err))
	}
}

func Release() {
	sqlDb, _ := db.DB()
	_ = sqlDb.Close()
}

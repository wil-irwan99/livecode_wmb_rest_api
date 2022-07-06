package manager

import (
	"log"
	"os"
	"wmb_rest_api/config"
	"wmb_rest_api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config *config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	env := os.Getenv("ENV")

	if env == "dev" {
		db = db.Debug()
	} else if env == "migration" {
		db = db.Debug()
		db.AutoMigrate(&model.Menu{}, &model.Table{}, &model.TransType{}, &model.Customer{}, &model.Discount{}, &model.MenuPrice{}, &model.Bill{}, &model.BillDetail{})
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

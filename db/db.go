package db

import (
	"ezsale/config"
	"ezsale/model"
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()

	sslConfig := "disable"
	if configuration.DB_SSL {
		sslConfig = `require`
	}

	connectString := fmt.Sprintf("host=%s port=%s sslmode=%s user=%s dbname=%s  password=%s ", configuration.DB_HOST, configuration.DB_PORT, sslConfig, configuration.DB_USERNAME, configuration.DB_NAME, configuration.DB_PASSWORD)
	fmt.Println(connectString)
	db, err = gorm.Open("postgres", connectString)
	// defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		panic("DB Connection Error")
	}
	db.AutoMigrate(
		&model.User{},
		&model.UserToken{},
		&model.ProductCategory{},
		&model.Product{},
		&model.UnitQuantity{},
		&model.Goods{},
	)

	db.LogMode(true)

	db.Callback().Create().Before("gorm:create").Register("delete_id_before_create", clearIDBerforeCreate)

}

func clearIDBerforeCreate(scope *gorm.Scope) {
	if scope.HasColumn("ID") {
		scope.SetColumn("ID", 0)
	}
}

func DbManager() *gorm.DB {
	return db
}

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
	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s ", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USERNAME, configuration.DB_NAME, configuration.DB_PASSWORD)
	fmt.Println(connectString)
	db, err = gorm.Open("postgres", connectString)
	// defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.User{}, &model.UserToken{})

}
func DbManager() *gorm.DB {
	return db
}

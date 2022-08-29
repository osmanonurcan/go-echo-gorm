package db

import (
	//"gorm.io/gorm"
	"fmt"
	//"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/osmanonurcan/go-test/config"
	"github.com/osmanonurcan/go-test/model"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	connect_string := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err = gorm.Open("mysql", connect_string)
	// defer db.Close()
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.Plan{}, &model.Student{})

}

func DbManager() *gorm.DB {
	return db
}

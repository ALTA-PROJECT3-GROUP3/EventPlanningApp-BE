package database

import (
	"fmt"
	"log"

	"github.com/ALTA-PROJECT3-GROUP3/EventPlanningApp-BE/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMySql(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER,
		cfg.DBPASSWORD,
		cfg.DBHOST,
		cfg.DBPORT,
		cfg.DBNAME,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to database:", err.Error())
		return nil
	}
	return db
}

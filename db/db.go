package db

import (
	"be-renthall/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	configuration := config.GetConfig()

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", configuration.DB_HOST, configuration.DB_PORT, configuration.DB_USERNAME, configuration.DB_NAME, configuration.DB_PASSWORD)
	db, err = gorm.Open("postgres", connectString)
	if err != nil {
		panic("DB1 Connection Error : " + err.Error() + " " + configuration.DB_HOST + " " + configuration.DB_PORT)
	}
}

func Manager() *gorm.DB {
	return db
}

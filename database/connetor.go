package database

import (
	"fmt"
	"maguas-blog-go/config"

	"github.com/jinzhu/gorm"
)

func Connect() (db *gorm.DB, err error) {
	connectConfig := config.DatabaseUser + ":" + config.DatabasePassword + "@/" + config.DatabaseName + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open("mysql", connectConfig)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return db, nil
}

package main

import (
	"fmt"
	"maguas-blog-go/model"

	"github.com/jinzhu/gorm"
)

var tables = []interface{}{
	&model.User{},
	&model.Article{},
	&model.Comment{},
	&model.Tag{},
}

func main() {

}

func DropAll(db *gorm.DB, t []interface{}) {
	var dropTables []interface{}
	if t != nil {
		dropTables = append(dropTables, t)
	} else {
		dropTables = append(dropTables, tables)
	}

	for _, table := range dropTables {
		tableName := db.NewScope(table).TableName()

		if db.HasTable(table) {
			db.DropTable(table)

			if !db.HasTable(table) {
				fmt.Printf("table %v drop successed\n", table)
			}
		} else {
			fmt.Printf("table %v inexistence\n", tableName)
		}
	}
}

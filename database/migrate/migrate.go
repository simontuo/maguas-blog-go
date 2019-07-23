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

func Migrate(db *gorm.DB, t interface{}) {
	var migrateTables []interface{}
	if t != nil {
		migrateTables = append(migrateTables, t)
	} else {
		migrateTables = append(migrateTables, tables)
	}

	for _, table := range migrateTables {
		tableName := db.NewScope(table).TableName()
		if db.HasTable(table) {
			fmt.Printf("table %v already migrated\n", tableName)
		} else {
			db.AutoMigrate(table)
			fmt.Printf("table %v migrate successed\n", tableName)
		}
	}
}

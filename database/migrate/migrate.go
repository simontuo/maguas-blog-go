package main

import (
	"flag"
	"fmt"
	"maguas-blog-go/database"
	"maguas-blog-go/model"
	"strings"

	"github.com/jinzhu/gorm"
)

var models = []interface{}{
	&model.User{},
	&model.Article{},
	&model.Comment{},
	&model.Tag{},
}

var all = flag.Bool("all", false, "migrate all tables")

var tables = flag.String("tables", "", "migrate table")

func main() {
	flag.Parse()

	db, _ := database.Connect()
	defer db.Close()

	if *all == true {
		Migrate(db, models)
	} else {
		if *tables == "" {
			fmt.Println("must specify tables")
			return
		}

		var migrateTables []interface{}

		for _, m := range models {
			tableName := db.NewScope(m).TableName()

			for _, t := range strings.Split(*tables, ",") {
				if tableName == t {
					migrateTables = append(migrateTables, m)
				}
			}
		}

		Migrate(db, migrateTables)
	}
}

func Migrate(db *gorm.DB, migrateModels []interface{}) {
	for _, table := range migrateModels {
		tableName := db.NewScope(table).TableName()
		if db.HasTable(table) {
			fmt.Printf("table %v already migrated\n", tableName)
		} else {
			db.AutoMigrate(table)
			fmt.Printf("table %v migrate successed\n", tableName)
		}
	}
}

func Table() {

}

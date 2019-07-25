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

var all = flag.Bool("all", false, "drop all tables")

var tables = flag.String("tables", "", "drop table")

func main() {
	flag.Parse()

	db, _ := database.Connect()
	defer db.Close()

	if *all {
		DropAll(db, models)
	} else {
		if *tables == "" {
			fmt.Println("must specify tables")
			return
		}

		for _, table := range strings.Split(*tables, ",") {
			if db.HasTable(table) {
				db.DropTable(table)

				if !db.HasTable(table) {
					fmt.Printf("table %v drop successed\n", table)
				}
			} else {
				fmt.Printf("table %v inexistence\n", table)
			}
		}
	}
}

func DropAll(db *gorm.DB, dropTables []interface{}) {
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

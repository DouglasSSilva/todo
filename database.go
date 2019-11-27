package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type todoModel struct {
	gorm.Model
	Title     string `db:"title" json:"title"`
	Completed int    `db:"completed" json:"completed"`
}

type transformedTodo struct {
	ID        uint   `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Completed bool   `db:"completed" json:"completed"`
}

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(172.17.0.2:3306)/todo?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

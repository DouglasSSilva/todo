package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// TodoModel struct to be inserted on db
type TodoModel struct {
	ID        uint   `db:"id" json:"ID"`
	Title     string `db:"title" json:"title"`
	Completed bool   `db:"completed" json:"completed"`
}

//TransformedTodo to be returned on get, will be removed

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "mysql", "root:root@tcp(172.17.0.2:3306)/todo?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&TodoModel{})
}

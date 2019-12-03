package app

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// TodoModel struct to be inserted on db
type TodoModel struct {
	// ID of the inserted todo
	ID uint `db:"id" json:"ID"`

	// Title defines the name of the to do to be inserted
	Title string `db:"title" json:"title"`

	// Completed  will be inserted on the db a true value tell the todo is finished
	Completed bool `db:"completed" json:"completed"`
}

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

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
    "models\"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	dsn := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable" // Update with your database credentials
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to database")
	}

	// Auto migrate the User model
	db.AutoMigrate(
		&Client{},
		&HistoryOrder{},
		&DepthOrder{},
	)
}



func main() {

}
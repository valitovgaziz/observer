package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/valitovgaziz/observer/model"
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

func getClients(c *gin.Context) {
	var clients []Client
	db.Find(&clients)
	c.JSON(200, gin.H{"data": clients})
}

func main() {
	r := gin.Default()

	r.GET("/clients", getClients)

	r.Run(":8282")
	fmt.Println("Listening on port 8282")
}

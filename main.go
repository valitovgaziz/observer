package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
		&DepthOrder{},
		&OrderBook{},
		&HistoryOrder{},
	)
}

type Client struct {
	Client_name   string `json:"client_name"`
	Exchange_name string `json:"exchange_name"`
	Label         string `json:"label"`
	Pair          string `json:"pair"`
}

type DepthOrder struct {
	Price   float64 `json:"price"`
	BaseQty float64 `json:"base_qty"`
}

type OrderBook struct {
	id       int64        `json:"id" gorm:"autoincrement"`
	exchange string       `json:"exchange"`
	Pair     string       `json:"pair"`
	asks     []DepthOrder `json:"asks"`
	bids     []DepthOrder `json:"bids"`
}

type HistoryOrder struct {
	Client_name           string    `json:"client_name"`
	Exchange_name         string    `json:"exchange_name"`
	Label                 string    `json:"label"`
	Pair                  string    `json:"pair"`
	Side                  string    `json:"side"`
	Type_                 string    `json:"type"`
	Base_qty              float64   `json:"base_qty"`
	Price                 float64   `json:"price"`
	Algorithm_name_placed string    `json:"algorithm_name_placed"`
	Lowest_sell_prc       float64   `json:"lowest_sell_prc"`
	Highest_buy_prc       float64   `json:"highest_buy_prc"`
	Commission_quote_qty  float64   `json:"commission_quote_qty"`
	Time_placed           time.Time `json:"time_placed"`
}

func SaveOrderBook(c *gin.Context) {
	var orderBook OrderBook
	c.BindJSON(&orderBook)
	eroor := db.Create(&orderBook)
	if eroor != nil {
		c.JSON(400, gin.H{"error": eroor})
		return
	}
	c.JSON(200, gin.H{"message": "Order book saved"})
}

func GetOrderBook(c *gin.Context) {
	var orders []HistoryOrder
	db.Find(&orders)
	c.JSON(200, gin.H{"data": orders})
}

func SaveOrder(c *gin.Context) {
	var order HistoryOrder
	c.BindJSON(&order)
	error := db.Create(&order)
	if error != nil {
		c.JSON(400, gin.H{"error": error})
		return
	}
	c.JSON(200, gin.H{"message": "Order saved"})
}

func GetOrderHistory(c *gin.Context) {
	var orders []HistoryOrder
	db.Find(&orders)
	c.JSON(200, gin.H{"data": orders})
}

func main() {
	r := gin.Default()

	r.GET("/GetOrderBook", GetOrderBook)
	r.POST("/SaveOrderBook", SaveOrderBook)
	r.GET("/GetOrderHistory", GetOrderHistory)
	r.POST("/GetOrderHistory", SaveOrder)

	r.Run(":8282")
	fmt.Println("Listening on port 8282")
}

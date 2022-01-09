package main

import (
	"bwastartup/handler"
	"bwastartup/stock"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	stockRepository := stock.NewRepository(db)
	stockService := stock.NewService(stockRepository)
	stockHandler := handler.NewStockHandler(stockService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/stocks", stockHandler.StockInput)

	router.Run(":3000")
}

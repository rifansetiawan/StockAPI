package stock

import "time"

type StockInput struct {
	Name        string    `json:"name" binding:"required"`
	Code        string    `json:"code" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Open        float64   `json:"open" binding:"required"`
	High        float64   `json:"high" binding:"required"`
	Low         float64   `json:"low" binding:"required"`
	Vol         float64   `json:"vol" binding:"required"`
	ChangeValue float64   `json:"change" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

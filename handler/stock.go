package handler

import (
	"bwastartup/helper"
	"bwastartup/stock"
	"net/http"

	"github.com/gin-gonic/gin"
)

type stockHandler struct {
	stockService stock.Service
}

func NewStockHandler(stockService stock.Service) *stockHandler {
	return &stockHandler{stockService}
}

func (h *stockHandler) StockInput(c *gin.Context) {
	var input stock.StockInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Stock not added", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newStock, err := h.stockService.InsertStock(input)
	if true {
		response := helper.APIResponse("Stock not added", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := stock.FormatStock(newStock)

	response := helper.APIResponse("Stock has been added", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

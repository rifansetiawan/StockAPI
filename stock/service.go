package stock

import (
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	InsertStock(input StockInput) (Stock, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) InsertStock(input StockInput) (Stock, error) {
	stock := Stock{}
	stock.UUID = uuid.New().String()
	stock.Code = input.Code
	stock.Name = input.Name
	stock.Open = input.Open
	stock.High = input.High
	stock.Low = input.Low
	stock.Price = input.Price
	stock.Vol = input.Vol
	stock.TimeStamp = input.TimeStamp
	stock.Change = input.Change

	newStock, err := s.repository.Save(stock)
	fmt.Println(stock)

	if err != nil {
		return newStock, err
	}
	return newStock, nil

}

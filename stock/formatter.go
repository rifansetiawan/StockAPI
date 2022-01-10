package stock

type StockFormatter struct {
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Open        float64 `json:"open"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Vol         string  `json:"vol"`
	Date        string  `json:"timestamp"`
	ChangeValue float64 `json:"change_value"`
}

func FormatStock(stock Stock) StockFormatter {
	formatter := StockFormatter{
		UUID:        stock.UUID,
		Name:        stock.Name,
		Code:        stock.Code,
		Price:       stock.Price,
		Open:        stock.Open,
		High:        stock.High,
		Low:         stock.Low,
		Vol:         stock.Vol,
		ChangeValue: stock.ChangeValue,
	}

	return formatter
}

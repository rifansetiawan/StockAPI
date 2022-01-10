package stock

import "time"

type Stock struct {
	UUID        string
	Name        string
	Code        string
	Price       float64
	Open        float64
	High        float64
	Low         float64
	Vol         string
	ChangeValue float64
	TimeStamp   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

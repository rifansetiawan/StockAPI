package stock

import "gorm.io/gorm"

type Repository interface {
	Save(stock Stock) (Stock, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(stock Stock) (Stock, error) {
	err := r.db.Create(&stock).Error
	if err != nil {
		return stock, err
	}
	return stock, nil
}

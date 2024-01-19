package service

import (
	"time"
)

type ProductService struct{}

type ProductRes struct {
	Title       string
	Description string
	Price       int32
}

func (s ProductService) Get(id int64) (*ProductRes, error) {
	if id != 1 {
		return nil, ErrRecordNotFound
	}

	time.Sleep(1 * time.Second)

	product := &ProductRes{
		Title:       "Fake Product",
		Description: "Fake Description",
		Price:       99,
	}

	return product, nil
}

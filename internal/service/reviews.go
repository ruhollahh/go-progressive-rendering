package service

import (
	"time"
)

type ReviewService struct{}

type ReviewForProductRes struct {
	Username string
	Content  string
	Date     time.Time
}

type ReviewsForProductRes []ReviewForProductRes

func (s ReviewService) GetAllForProduct(productID int64) (ReviewsForProductRes, error) {
	if productID != 1 {
		return nil, nil
	}

	time.Sleep(2 * time.Second)

	reviews := ReviewsForProductRes{
		{
			Username: "Fake User 1",
			Content:  "Fake Review 1",
			Date:     time.Now(),
		},
		{
			Username: "Fake User 2",
			Content:  "Fake Review 2",
			Date:     time.Now(),
		},
		{
			Username: "Fake User 3",
			Content:  "Fake Review 3",
			Date:     time.Now(),
		},
	}

	return reviews, nil
}

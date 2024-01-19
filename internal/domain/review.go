package domain

import "time"

type Review struct {
	ID        int64
	ProductID int64
	UserID    int64
	Content   string
	CreatedAt time.Time
}

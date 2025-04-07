package entity

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusDelivery  Status = "delivery"
	StatusCompleted Status = "completed"
	StatusCanceled  Status = "canceled"
)

type Order struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    Status
}

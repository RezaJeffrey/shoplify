package domain

import "time"

type Order struct {
	ID     int64
	Number string
	Status OrderStatus

	Items []OrderItem

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderStatus uint8

const (
	Submitted OrderStatus = iota + 1
	Approved
	Failed
)

var statusTitle = map[OrderStatus]string{
	Submitted: "Submitted",
	Approved:  "Approved",
	Failed:    "Failed",
}

func (os OrderStatus) String() string {
	return statusTitle[os]
}

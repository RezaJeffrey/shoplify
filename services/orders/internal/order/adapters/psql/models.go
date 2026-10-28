package psql

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderModel struct {
	ID     int64
	Number string
	Status uint8

	Items []OrderItemModel `gorm:"foreignKey:OrderID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItemModel struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  decimal.Decimal
	UnitPrice decimal.Decimal
}

func (OrderModel) TableName() string {
	return "orders"
}

func (OrderItemModel) TableName() string {
	return "order_items"
}

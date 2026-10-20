package domain

import "github.com/shopspring/decimal"

type OrderItem struct {
	ID        int64
	OrderID   int64
	ProductID int64
	Quantity  decimal.Decimal
	UnitPrice decimal.Decimal
}

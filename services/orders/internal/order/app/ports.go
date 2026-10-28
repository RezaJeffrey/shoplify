package app

import (
	"context"

	"github.com/rezajeffrey/shoplify/services/orders/internal/order/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, o *domain.Order) error
	FetchByID(ctx context.Context, id int64) (*domain.Order, error)
}

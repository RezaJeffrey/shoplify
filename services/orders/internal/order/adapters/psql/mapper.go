package psql

import (
	"fmt"

	"github.com/rezajeffrey/shoplify/services/orders/internal/order/domain"
)

func ToPersistenceModel(do *domain.Order) (*OrderModel, error) {
	if do.Status.Valid() {
		return nil, fmt.Errorf("invalid order status to store: %d", uint8(do.Status))
	}

	model := &OrderModel{
		ID:        do.ID,
		Number:    do.Number,
		Status:    uint8(do.Status),
		CreatedAt: do.CreatedAt,
		UpdatedAt: do.UpdatedAt,
	}

	for _, i := range do.Items {
		item := OrderItemModel{
			ID:        i.ID,
			OrderID:   i.OrderID,
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			UnitPrice: i.UnitPrice,
		}

		model.Items = append(model.Items, item)
	}

	return model, nil
}

func ToDomainModel(po *OrderModel) (*domain.Order, error) {
	if err := validateOrderStatus(domain.OrderStatus(po.Status)); err != nil {
		return nil, err
	}

	model := &domain.Order{
		ID:        po.ID,
		Number:    po.Number,
		Status:    domain.OrderStatus(po.Status),
		CreatedAt: po.CreatedAt,
		UpdatedAt: po.UpdatedAt,
	}

	for _, i := range po.Items {
		item := domain.OrderItem{
			ID:        i.ID,
			OrderID:   i.OrderID,
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			UnitPrice: i.UnitPrice,
		}

		model.Items = append(model.Items, item)
	}

	return model, nil
}

func validateOrderStatus(status domain.OrderStatus) error {
	switch status {
	case domain.Approved, domain.Failed, domain.Submitted:
	default:
		return fmt.Errorf("unkown domain status: %d", uint8(status))
	}
	return nil
}

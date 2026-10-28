package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/rezajeffrey/shoplify/services/orders/internal/order/domain"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func (r *OrderRepository) Create(ctx context.Context, o *domain.Order) error {
	model, err := ToPersistenceModel(o)
	if err != nil {
		return err
	}

	if err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Items").Create(model).Error; err != nil {
			return err
		}

		if len(model.Items) > 0 {
			if err := tx.Create(&model.Items).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) FetchByID(ctx context.Context, ID int64) (*domain.Order, error) {
	var model *OrderModel

	err := r.db.WithContext(ctx).
		Preload("Items").
		Where("id = ?", ID).
		First(&model).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("Order with id: %d not found: %w", ID, err)
	}

	if err != nil {
		return nil, fmt.Errorf("Error while trying to fetch order with id %d: %w", ID, err)
	}

	return ToDomainModel(model)
}

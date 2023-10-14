package local

import (
	"errors"
	"fmt"
	"restaurant-api/internal/support/orders/core/entity"
)

type Repository struct {
	orders []*entity.Order
}

func NewRepository() Repository {
	return Repository{
		orders: []*entity.Order{},
	}
}

func (r *Repository) CreateOrder(order *entity.Order) (*entity.Order, error) {
	index := len(r.orders)
	order.ID = index + 1

	r.orders = append(r.orders, order)
	return order, nil
}

func (r *Repository) GetOrderIndex(orderID int) (int, error) {
	for i, order := range r.orders {
		if order.ID == orderID {
			return i, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("order not found with id: %d", orderID))
}

func (r *Repository) UpdateOrder(orderID int, order entity.Order) error {
	i, err := r.GetOrderIndex(orderID)
	if err != nil {
		return err
	}

	r.orders[i].UpdatedAt = order.UpdatedAt
	r.orders[i].Status = order.Status

	return nil
}

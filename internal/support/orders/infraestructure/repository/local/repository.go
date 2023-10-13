package local

import "restaurant-api/internal/support/orders/core/entity"

type Repository struct {
	orders []entity.Order
}

func NewRepository() Repository {
	return Repository{
		orders: []entity.Order{},
	}
}

func (r *Repository) CreateOrder(order *entity.Order) (*entity.Order, error) {
	index := len(r.orders)
	order.ID = index + 1

	r.orders = append(r.orders, *order)
	return order, nil
}

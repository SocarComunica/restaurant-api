package mapper

import (
	"restaurant-api/internal/support/orders/core/entity"
	"restaurant-api/src/api/handler/order/contract"
)

type Mapper struct{}

func (m Mapper) MapEntityToResponse(order entity.Order) contract.Order {
	return contract.Order{
		ID:        order.ID,
		DishID:    order.DishID,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Status:    order.Status,
	}
}

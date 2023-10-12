package mapper

import (
	"restaurant-api/internal/support/dishes/core/entity"
	"restaurant-api/src/api/handler/dish/contract"
)

type Mapper struct{}

func (m Mapper) MapEntityToResponse(dish entity.Dish) contract.Dish {
	return contract.Dish{
		ID:     dish.ID,
		Name:   dish.Name,
		Recipe: dish.Recipe,
		Stats: contract.Stats{
			Finished: dish.Stats.Finished,
			Pending:  dish.Stats.Pending,
		},
	}
}

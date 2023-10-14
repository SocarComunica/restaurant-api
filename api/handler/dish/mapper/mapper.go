package mapper

import (
	contract2 "restaurant-api/api/handler/dish/contract"
	"restaurant-api/internal/support/dishes/core/entity"
)

type Mapper struct{}

func (m Mapper) MapEntityToResponse(dish entity.Dish) contract2.Dish {
	return contract2.Dish{
		ID:     dish.ID,
		Name:   dish.Name,
		Recipe: dish.Recipe,
		Stats: contract2.Stats{
			Delivered:  dish.Stats.Delivered,
			Finished:   dish.Stats.Finished,
			InProgress: dish.Stats.InProgress,
			Queued:     dish.Stats.Queued,
		},
	}
}

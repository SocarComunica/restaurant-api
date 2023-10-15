package local

import (
	"errors"
	"fmt"
	"restaurant-api/internal/support/dishes/core/entity"
)

type Repository struct {
	dishes []*entity.Dish
}

func NewRepository() Repository {
	return Repository{
		dishes: loadDishes(),
	}
}

func (r *Repository) GetAllDishes() ([]entity.Dish, error) {
	dishes := make([]entity.Dish, len(r.dishes))
	for i, dish := range r.dishes {
		dishes[i] = *dish
	}
	return dishes, nil
}

func (r *Repository) GetDish(dishID int) (entity.Dish, error) {
	i, err := r.GetDishIndex(dishID)
	if err != nil {
		return entity.Dish{}, err
	}

	return *r.dishes[i], nil
}

func (r *Repository) GetDishIndex(dishID int) (int, error) {
	for i, dish := range r.dishes {
		if dish.ID == dishID {
			return i, nil
		}
	}

	return 0, errors.New(fmt.Sprintf("dish not found with id: %d", dishID))
}

func (r *Repository) UpdateDish(dishID int, dish entity.Dish) error {
	i, err := r.GetDishIndex(dishID)
	if err != nil {
		return err
	}

	r.dishes[i].Stats.Queued = dish.Stats.Queued
	r.dishes[i].Stats.InProgress = dish.Stats.InProgress
	r.dishes[i].Stats.Finished = dish.Stats.Finished
	r.dishes[i].Stats.Delivered = dish.Stats.Delivered

	return nil
}

func loadDishes() []*entity.Dish {
	dishes := []*entity.Dish{
		{
			ID:   1,
			Name: "Ensalada de Pollo con Limón y Queso",
			Recipe: map[string]int{
				"chicken": 2,
				"lemon":   1,
				"lettuce": 1,
				"cheese":  1, // Tratar el queso como una rebanada
				"onion":   1,
				"tomato":  1,
			},
		},
		{
			ID:   2,
			Name: "Arroz con Tomate y Pollo",
			Recipe: map[string]int{
				"chicken": 2,
				"tomato":  2,
				"rice":    1,
				"onion":   1,
				"lemon":   1,
				"ketchup": 2,
			},
		},
		{
			ID:   3,
			Name: "Papas a la Francesa con Queso",
			Recipe: map[string]int{
				"potato":  4,
				"cheese":  1, // Tratar el queso como una rebanada
				"ketchup": 2,
			},
		},
		{
			ID:   4,
			Name: "Hamburguesas de Carne con Cebolla",
			Recipe: map[string]int{
				"meat":    1, // Tratar la carne como 1 libra
				"onion":   1,
				"tomato":  1,
				"lettuce": 2,
				"cheese":  1, // Tratar el queso como una rebanada
				"ketchup": 2,
			},
		},
		{
			ID:   5,
			Name: "Pollo al Limón con Arroz",
			Recipe: map[string]int{
				"chicken": 2,
				"lemon":   2,
				"rice":    1,
				"onion":   1,
				"ketchup": 2,
			},
		},
		{
			ID:   6,
			Name: "Arroz con Pollo y Tomate",
			Recipe: map[string]int{
				"chicken": 2,
				"tomato":  2,
				"rice":    1,
				"onion":   1,
				"ketchup": 2,
			},
		},
	}

	return dishes
}

package local

import (
	"errors"
	"restaurant-api/internal/support/dishes/core/entity"
)

type LocalRepository struct {
	dishes []entity.Dish
}

func NewLocalRepository() LocalRepository {
	return LocalRepository{
		dishes: loadDishes(),
	}
}

func (lr LocalRepository) GetAllDishes() ([]entity.Dish, error) {
	return lr.dishes, nil
}

func (lr LocalRepository) GetDish(dishID int) (entity.Dish, error) {
	for _, dish := range lr.dishes {
		if dish.ID == dishID {
			return dish, nil
		}
	}

	return entity.Dish{}, errors.New("dish not found with id")
}

func (lr LocalRepository) UpdateDish(dishID int, dish entity.Dish) error {
	panic("implement update dish method on repository")
}

func loadDishes() []entity.Dish {
	dishes := []entity.Dish{
		{
			ID:   1,
			Name: "Ensalada de Pollo con Limón y Queso",
			Recipe: map[string]int{
				"Chicken": 2,
				"Lemon":   1,
				"Lettuce": 1,
				"Cheese":  1, // Tratar el queso como una rebanada
				"Onion":   1,
				"Tomato":  1,
			},
		},
		{
			ID:   2,
			Name: "Arroz con Tomate y Pollo",
			Recipe: map[string]int{
				"Chicken": 2,
				"Tomato":  2,
				"Rice":    1,
				"Onion":   1,
				"Lemon":   1,
				"Ketchup": 2,
			},
		},
		{
			ID:   3,
			Name: "Papas a la Francesa con Queso",
			Recipe: map[string]int{
				"Potato":  4,
				"Cheese":  1, // Tratar el queso como una rebanada
				"Ketchup": 2,
			},
		},
		{
			ID:   4,
			Name: "Hamburguesas de Carne con Cebolla",
			Recipe: map[string]int{
				"Meat":    1, // Tratar la carne como 1 libra
				"Onion":   1,
				"Tomato":  1,
				"Lettuce": 2,
				"Cheese":  1, // Tratar el queso como una rebanada
				"Ketchup": 2,
			},
		},
		{
			ID:   5,
			Name: "Pollo al Limón con Arroz",
			Recipe: map[string]int{
				"Chicken": 2,
				"Lemon":   2,
				"Rice":    1,
				"Onion":   1,
				"Ketchup": 2,
			},
		},
		{
			ID:   6,
			Name: "Arroz con Pollo y Tomate",
			Recipe: map[string]int{
				"Chicken": 2,
				"Tomato":  2,
				"Rice":    1,
				"Onion":   1,
				"Ketchup": 2,
			},
		},
	}

	return dishes
}

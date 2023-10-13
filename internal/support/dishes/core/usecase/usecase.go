package usecase

import (
	"github.com/gin-gonic/gin"
	"restaurant-api/internal/support/dishes/core/entity"
)

type Repository interface {
	GetAllDishes() ([]entity.Dish, error)
	GetDish(dishID int) (entity.Dish, error)
	UpdateDish(dishID int, dish entity.Dish) error
}

type UseCase struct {
	repository Repository
}

func NewDishesUseCase(repository Repository) UseCase {
	return UseCase{
		repository: repository,
	}
}

func (u *UseCase) GetAllDishes(ctx *gin.Context) ([]entity.Dish, error) {
	return u.repository.GetAllDishes()
}

func (u *UseCase) GetDish(ctx *gin.Context, dishID int) (*entity.Dish, error) {
	dish, err := u.repository.GetDish(dishID)
	if err != nil {
		return nil, err
	}

	return &dish, nil
}

func (u *UseCase) UpdateDishStats(ctx *gin.Context, dishID int, stats entity.Stats) error {
	dish, err := u.GetDish(ctx, dishID)
	if err != nil {
		return nil
	}

	dish.Stats = stats
	return u.repository.UpdateDish(dishID, *dish)
}

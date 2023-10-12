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

func (u UseCase) GetAllDishes(ctx *gin.Context) ([]entity.Dish, error) {
	return u.repository.GetAllDishes()
}

func (u UseCase) GetDish(ctx *gin.Context, dishID int) (entity.Dish, error) {
	panic("implement GetDish on usecase")
}

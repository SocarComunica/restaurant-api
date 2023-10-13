package usecase

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	dishesEntity "restaurant-api/internal/support/dishes/core/entity"
	"restaurant-api/internal/support/orders/core/entity"
	"restaurant-api/internal/support/orders/infraestructure/queue"
	"time"
)

type Repository interface {
	CreateOrder(order *entity.Order) (*entity.Order, error)
}

type DishesUseCase interface {
	GetAllDishes(ctx *gin.Context) ([]dishesEntity.Dish, error)
	UpdateDishStats(ctx *gin.Context, dishID int, stats dishesEntity.Stats) error
}

type UseCase struct {
	repository    Repository
	dishesUseCase DishesUseCase
	queue         *queue.Queue
}

func NewOrdersUseCase(dishesUseCase DishesUseCase, repository Repository) UseCase {
	return UseCase{
		repository:    repository,
		dishesUseCase: dishesUseCase,
		queue:         &queue.Queue{},
	}
}

func (u *UseCase) NewRandomOrder(ctx *gin.Context) (*entity.Order, error) {
	dishes, err := u.dishesUseCase.GetAllDishes(ctx)
	if err != nil {
		return nil, err
	}

	index := rand.Intn(len(dishes))
	dish := dishes[index]

	order := entity.Order{
		DishID:    dish.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "Queued",
	}

	orderWithId, err := u.repository.CreateOrder(&order)
	if err != nil {
		return nil, err
	}

	dish.Stats.Queued++
	err = u.dishesUseCase.UpdateDishStats(ctx, dish.ID, dish.Stats)
	if err != nil {
		return nil, err
	}
	u.queue.Enqueue(*orderWithId)

	return orderWithId, nil
}

func (u *UseCase) GetQueue(ctx *gin.Context) []entity.Order {
	var q []entity.Order
	for _, order := range *u.queue {
		q = append(q, *order)
	}
	return q
}

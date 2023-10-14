package usecase

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	dishesEntity "restaurant-api/internal/support/dishes/core/entity"
	"restaurant-api/internal/support/orders/core/entity"
	"restaurant-api/internal/support/orders/infraestructure/queue"
	"time"
)

const (
	QueuedStatus     = "queued"
	InProgressStatus = "in_progress"
	FinishedStatus   = "finished"
	DeliveredStatus  = "delivered"
)

type Repository interface {
	CreateOrder(order *entity.Order) (*entity.Order, error)
	UpdateOrder(orderID int, order entity.Order) error
}

type DishesUseCase interface {
	GetAllDishes(ctx *gin.Context) ([]dishesEntity.Dish, error)
	GetDish(ctx *gin.Context, dishID int) (*dishesEntity.Dish, error)
	UpdateDishStats(ctx *gin.Context, dishID int, stats dishesEntity.Stats) error
}

type UseCase struct {
	repository       Repository
	dishesUseCase    DishesUseCase
	queuedOrders     *queue.Queue
	inProgressOrders *queue.Queue
	finishedOrders   *queue.Queue
}

func NewOrdersUseCase(dishesUseCase DishesUseCase, repository Repository) UseCase {
	return UseCase{
		repository:       repository,
		dishesUseCase:    dishesUseCase,
		queuedOrders:     &queue.Queue{},
		inProgressOrders: &queue.Queue{},
		finishedOrders:   &queue.Queue{},
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
		Status:    QueuedStatus,
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
	u.queuedOrders.Enqueue(*orderWithId)

	return orderWithId, nil
}

func (u *UseCase) UpdateOrderQueuedToInProgress(ctx *gin.Context) error {
	order, len0 := u.queuedOrders.Dequeue()
	if len0 {
		return errors.New("there is no more queued orders")
	}

	if order == nil {
		return errors.New("next queued order is nil")
	}

	order.Status = InProgressStatus
	order.UpdatedAt = time.Now()

	err := u.repository.UpdateOrder(order.ID, *order)
	if err != nil {
		return err
	}

	dish, err := u.dishesUseCase.GetDish(ctx, order.DishID)
	if err != nil {
		return err
	}
	if dish == nil {
		return errors.New(fmt.Sprintf("dish with ID: %d is not present", order.DishID))
	}
	dish.Stats.Queued--
	dish.Stats.InProgress++
	err = u.dishesUseCase.UpdateDishStats(ctx, dish.ID, dish.Stats)
	if err != nil {
		return err
	}

	u.inProgressOrders.Enqueue(*order)
	return nil
}

func (u *UseCase) UpdateOrderInProgressToFinished(ctx *gin.Context) error {
	order, len0 := u.inProgressOrders.Dequeue()
	if len0 {
		return errors.New("there is no more in progress orders")
	}

	if order == nil {
		return errors.New("next in progress order is nil")
	}

	order.Status = FinishedStatus
	order.UpdatedAt = time.Now()

	err := u.repository.UpdateOrder(order.ID, *order)
	if err != nil {
		return err
	}

	dish, err := u.dishesUseCase.GetDish(ctx, order.DishID)
	if err != nil {
		return err
	}
	if dish == nil {
		return errors.New(fmt.Sprintf("dish with ID: %d is not present", order.DishID))
	}
	dish.Stats.InProgress--
	dish.Stats.Finished++
	err = u.dishesUseCase.UpdateDishStats(ctx, dish.ID, dish.Stats)
	if err != nil {
		return err
	}

	u.finishedOrders.Enqueue(*order)
	return nil
}

func (u *UseCase) UpdateOrderFinishedToDelivered(ctx *gin.Context) error {
	order, len0 := u.queuedOrders.Dequeue()
	if len0 {
		return errors.New("there is no more finished orders")
	}

	if order == nil {
		return errors.New("next finished order is nil")
	}

	order.Status = DeliveredStatus
	order.UpdatedAt = time.Now()

	err := u.repository.UpdateOrder(order.ID, *order)
	if err != nil {
		return err
	}

	dish, err := u.dishesUseCase.GetDish(ctx, order.DishID)
	if err != nil {
		return err
	}
	if dish == nil {
		return errors.New(fmt.Sprintf("dish with ID: %d is not present", order.DishID))
	}
	dish.Stats.Finished--
	dish.Stats.Delivered++
	err = u.dishesUseCase.UpdateDishStats(ctx, dish.ID, dish.Stats)
	if err != nil {
		return err
	}

	u.inProgressOrders.Enqueue(*order)
	return nil
}

func (u *UseCase) GetQueuedOrdersQueue(ctx *gin.Context) []entity.Order {
	return getQueue(u.queuedOrders)
}

func (u *UseCase) GetInProgressOrdersQueue(ctx *gin.Context) []entity.Order {
	return getQueue(u.inProgressOrders)
}

func (u *UseCase) GetFinishedOrdersQueue(ctx *gin.Context) []entity.Order {
	return getQueue(u.finishedOrders)
}

func getQueue(q *queue.Queue) []entity.Order {
	var result []entity.Order
	for _, order := range *q {
		result = append(result, *order)
	}
	return result
}

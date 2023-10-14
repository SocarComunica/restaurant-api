package dependence

import (
	"restaurant-api/src/api/dependence/container"
	"restaurant-api/src/api/handler"
	dishesHandler "restaurant-api/src/api/handler/dish"
	dishesMapper "restaurant-api/src/api/handler/dish/mapper"
	ordersHandler "restaurant-api/src/api/handler/order"
	ordersMapper "restaurant-api/src/api/handler/order/mapper"
)

type HandlerContainer struct {
	container                       container.Container
	GetAllDishHandler               handler.Handler
	GetDishHandler                  handler.Handler
	NewRandomOrderHandler           handler.Handler
	GetQueuedOrdersQueueHandler     handler.Handler
	GetInProgressOrdersQueueHandler handler.Handler
	GetFinishedOrdersQueueHandler   handler.Handler
}

type StartApp struct {
	container        container.Container
	useCaseContainer UseCaseContainer
}

func NewWire() HandlerContainer {
	c := container.NewContainer()
	u := NewUseCase()
	startApp := StartApp{c, u}
	return HandlerContainer{
		container:                       c,
		GetAllDishHandler:               startApp.NewGetAllDishHandler(),
		GetDishHandler:                  startApp.NewGetDishHandler(),
		NewRandomOrderHandler:           startApp.NewNewRandomOrderHandler(),
		GetQueuedOrdersQueueHandler:     startApp.NewGetQueuedOrdersQueueHandler(),
		GetInProgressOrdersQueueHandler: startApp.NewGetInProgressOrdersQueueHandler(),
		GetFinishedOrdersQueueHandler:   startApp.NewGetFinishedOrdersQueueHandler(),
	}
}

func (s StartApp) NewGetAllDishHandler() handler.Handler {
	return dishesHandler.NewGetAllDishHandler(&s.useCaseContainer.DishesUseCase, dishesMapper.Mapper{})
}

func (s StartApp) NewGetDishHandler() handler.Handler {
	return dishesHandler.NewGetHandler(&s.useCaseContainer.DishesUseCase, dishesMapper.Mapper{})
}

func (s StartApp) NewNewRandomOrderHandler() handler.Handler {
	return ordersHandler.NewNewRandomOrderHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewGetQueuedOrdersQueueHandler() handler.Handler {
	return ordersHandler.NewGetQueuedOrdersQueueHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewGetInProgressOrdersQueueHandler() handler.Handler {
	return ordersHandler.NewGetInProgressOrdersQueueHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewGetFinishedOrdersQueueHandler() handler.Handler {
	return ordersHandler.NewGetFinishedOrdersQueueHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

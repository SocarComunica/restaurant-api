package dependence

import (
	"restaurant-api/api/dependence/container"
	"restaurant-api/api/handler"
	"restaurant-api/api/handler/dish"
	dishesMapper "restaurant-api/api/handler/dish/mapper"
	"restaurant-api/api/handler/order"
	ordersMapper "restaurant-api/api/handler/order/mapper"
)

type HandlerContainer struct {
	container                       container.Container
	GetAllDishHandler               handler.Handler
	GetDishHandler                  handler.Handler
	CreateNewRandomOrderHandler     handler.Handler
	GetQueuedOrdersQueueHandler     handler.Handler
	GetInProgressOrdersQueueHandler handler.Handler
	GetFinishedOrdersQueueHandler   handler.Handler
	UpdateOrdersQueuesHandler       handler.Handler
}

type StartApp struct {
	container        container.Container
	useCaseContainer UseCaseContainer
}

func NewWire() HandlerContainer {
	c := container.NewContainer()
	u := NewUseCase(c)
	startApp := StartApp{c, u}
	return HandlerContainer{
		container:                       c,
		GetAllDishHandler:               startApp.NewGetAllDishHandler(),
		GetDishHandler:                  startApp.NewGetDishHandler(),
		CreateNewRandomOrderHandler:     startApp.NewNewRandomOrderHandler(),
		GetQueuedOrdersQueueHandler:     startApp.NewGetQueuedOrdersQueueHandler(),
		GetInProgressOrdersQueueHandler: startApp.NewGetInProgressOrdersQueueHandler(),
		GetFinishedOrdersQueueHandler:   startApp.NewGetFinishedOrdersQueueHandler(),
		UpdateOrdersQueuesHandler:       startApp.NewUpdateOrdersQueuesHandler(),
	}
}

func (s StartApp) NewGetAllDishHandler() handler.Handler {
	return dish.NewGetAllDishHandler(&s.useCaseContainer.DishesUseCase, dishesMapper.Mapper{})
}

func (s StartApp) NewGetDishHandler() handler.Handler {
	return dish.NewGetHandler(&s.useCaseContainer.DishesUseCase, dishesMapper.Mapper{})
}

func (s StartApp) NewNewRandomOrderHandler() handler.Handler {
	return order.NewNewRandomOrderHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewGetQueuedOrdersQueueHandler() handler.Handler {
	return order.NewGetQueuedOrdersQueueHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewGetInProgressOrdersQueueHandler() handler.Handler {
	return order.NewGetInProgressOrdersQueueHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewGetFinishedOrdersQueueHandler() handler.Handler {
	return order.NewGetFinishedOrdersQueueHandler(&s.useCaseContainer.OrdersUseCase, ordersMapper.Mapper{})
}

func (s StartApp) NewUpdateOrdersQueuesHandler() handler.Handler {
	return order.NewUpdateOrdersQueuesHandler(&s.useCaseContainer.OrdersUseCase)
}

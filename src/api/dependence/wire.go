package dependence

import (
	"restaurant-api/src/api/dependence/container"
	"restaurant-api/src/api/handler"
	dishesHandler "restaurant-api/src/api/handler/dish"
	dishesMapper "restaurant-api/src/api/handler/dish/mapper"
)

type HandlerContainer struct {
	container         container.Container
	GetAllDishHandler handler.Handler
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
		container:         c,
		GetAllDishHandler: startApp.newGetAllDishHandler(),
	}
}

func (s *StartApp) newGetAllDishHandler() handler.Handler {
	return dishesHandler.NewGetAllDishHandler(s.useCaseContainer.GetDishesUseCase(), dishesMapper.Mapper{})
}

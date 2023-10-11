package dependence

import (
	dishUseCase "restaurant-api/internal/support/dishes/core/usecase"
	"restaurant-api/src/api/dependence/container"
	"restaurant-api/src/api/handler"
	"restaurant-api/src/api/handler/dish"
)

type HandlerContainer struct {
	container   container.Container
	DishHandler handler.Handler
}

type StartApp struct {
	container container.Container
}

func NewWire() HandlerContainer {
	c := container.NewContainer()
	startApp := StartApp{c}
	return HandlerContainer{
		container:   c,
		DishHandler: startApp.newDishHandler(),
	}
}

func (s *StartApp) newDishHandler() handler.Handler {
	// TODO: Implement use case
	useCase := dishUseCase.NewRandomDishUseCase()
	return dish.NewDishHandler(useCase)
}

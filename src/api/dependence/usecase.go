package dependence

import (
	dishesUseCase "restaurant-api/internal/support/dishes/core/usecase"
	localDishesRepository "restaurant-api/internal/support/dishes/infraestructure/repository/local"
	"restaurant-api/src/api/dependence/container"
)

type UseCaseContainer struct {
	container     container.Container
	DishesUseCase *dishesUseCase.UseCase
}

func NewUseCase(c container.Container) UseCaseContainer {
	return UseCaseContainer{
		container: c,
	}
}

func (u UseCaseContainer) GetDishesUseCase() dishesUseCase.UseCase {
	if u.DishesUseCase == nil {
		repository := localDishesRepository.NewLocalRepository()
		useCase := dishesUseCase.NewDishesUseCase(repository)
		u.DishesUseCase = &useCase
	}
	return *u.DishesUseCase
}

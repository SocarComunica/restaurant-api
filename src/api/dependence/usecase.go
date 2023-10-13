package dependence

import (
	dishesUseCase "restaurant-api/internal/support/dishes/core/usecase"
	localDishesRepository "restaurant-api/internal/support/dishes/infraestructure/repository/local"
	ordersUseCase "restaurant-api/internal/support/orders/core/usecase"
	localOrdersUseCase "restaurant-api/internal/support/orders/infraestructure/repository/local"
)

var (
	dishes *dishesUseCase.UseCase
	orders *ordersUseCase.UseCase
)

type UseCaseContainer struct {
	DishesUseCase dishesUseCase.UseCase
	OrdersUseCase ordersUseCase.UseCase
}

func NewUseCase() UseCaseContainer {
	dishes = GetDishesUseCase()
	orders = GetOrdersUseCase()
	return UseCaseContainer{
		DishesUseCase: *dishes,
		OrdersUseCase: *orders,
	}
}

func GetDishesUseCase() *dishesUseCase.UseCase {
	if dishes == nil {
		repository := localDishesRepository.NewRepository()
		useCase := dishesUseCase.NewDishesUseCase(&repository)
		dishes = &useCase
	}
	return dishes
}

func GetOrdersUseCase() *ordersUseCase.UseCase {
	if orders == nil {
		repository := localOrdersUseCase.NewRepository()
		useCase := ordersUseCase.NewOrdersUseCase(GetDishesUseCase(), &repository)
		orders = &useCase
	}
	return orders
}

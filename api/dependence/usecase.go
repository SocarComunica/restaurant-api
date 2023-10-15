package dependence

import (
	"restaurant-api/api/dependence/container"
	dishesUseCase "restaurant-api/internal/support/dishes/core/usecase"
	localDishesRepository "restaurant-api/internal/support/dishes/infraestructure/repository/local"
	ordersUseCase "restaurant-api/internal/support/orders/core/usecase"
	localOrdersUseCase "restaurant-api/internal/support/orders/infraestructure/repository/local"
	warehouseUseCase "restaurant-api/internal/support/warehouse/core/usecase"
	localWarehouseRepository "restaurant-api/internal/support/warehouse/infraestructure/repository/local"
)

var (
	dishes    *dishesUseCase.UseCase
	orders    *ordersUseCase.UseCase
	warehouse *warehouseUseCase.UseCase
)

type UseCaseContainer struct {
	c                container.Container
	DishesUseCase    dishesUseCase.UseCase
	OrdersUseCase    ordersUseCase.UseCase
	WarehouseUseCase warehouseUseCase.UseCase
}

func NewUseCase(c container.Container) UseCaseContainer {
	dishes = GetDishesUseCase()
	orders = GetOrdersUseCase(c)
	warehouse = GetWarehouseUseCase(c)
	return UseCaseContainer{
		c:                c,
		DishesUseCase:    *dishes,
		OrdersUseCase:    *orders,
		WarehouseUseCase: *warehouse,
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

func GetWarehouseUseCase(c container.Container) *warehouseUseCase.UseCase {
	if warehouse == nil {
		repository := localWarehouseRepository.NewRepository()
		useCase := warehouseUseCase.NewWarehouseUseCase(&repository, c.RestClient.Market)
		warehouse = &useCase
	}
	return warehouse
}

func GetOrdersUseCase(c container.Container) *ordersUseCase.UseCase {
	if orders == nil {
		repository := localOrdersUseCase.NewRepository()
		useCase := ordersUseCase.NewOrdersUseCase(&repository, GetDishesUseCase(), GetWarehouseUseCase(c))
		orders = &useCase
	}
	return orders
}

package dish

import (
	"github.com/gin-gonic/gin"
	"net/http"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/dishes/core/entity"
	"restaurant-api/src/api/errors"
	"restaurant-api/src/api/handler/dish/contract"
	"restaurant-api/src/api/handler/dish/mapper"
)

type GetAllDishesUseCase interface {
	GetAllDishes(ctx *gin.Context) ([]entity.Dish, error)
}

type Mapper interface {
	MapEntityToResponse(dish entity.Dish) contract.Dish
}

type GetAllHandler struct {
	useCase GetAllDishesUseCase
	mapper  Mapper
}

func NewGetAllDishHandler(useCase GetAllDishesUseCase, mapper mapper.Mapper) GetAllHandler {
	return GetAllHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h GetAllHandler) Handler(ginCtx *gin.Context) {
	errors.ErrorWrapper(h.handler, ginCtx)
}

func (h GetAllHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	dishes, err := h.useCase.GetAllDishes(ginCtx)
	if err != nil {
		apiErr := platformerrors.NewInternalServerAPIError("An error has occurred while processing your request", err)
		return &apiErr
	}

	var result []contract.Dish
	for _, dish := range dishes {
		result = append(result, h.mapper.MapEntityToResponse(dish))
	}

	ginCtx.JSON(http.StatusOK, result)
	return nil
}

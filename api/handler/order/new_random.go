package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	commonErrors "restaurant-api/api/errors"
	"restaurant-api/api/handler/order/contract"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/orders/core/entity"
)

const (
	NotFoundDishIDParamErrorMessage = "there is no resource available with this id"
)

type NewRandomOrderUseCase interface {
	NewRandomOrder(ctx *gin.Context) (*entity.Order, error)
}

type Mapper interface {
	MapEntityToResponse(order entity.Order) contract.Order
}

type NewRandomOrderHandler struct {
	useCase NewRandomOrderUseCase
	mapper  Mapper
}

func NewNewRandomOrderHandler(useCase NewRandomOrderUseCase, mapper Mapper) NewRandomOrderHandler {
	return NewRandomOrderHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h NewRandomOrderHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h NewRandomOrderHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	order, err := h.useCase.NewRandomOrder(ginCtx)
	var apiErr platformerrors.APIError
	if err != nil {
		if order == nil {
			apiErr = platformerrors.NewNotFoundAPIError(NotFoundDishIDParamErrorMessage, err)
		} else {
			apiErr = platformerrors.NewInternalServerAPIError("An error has occurred while processing your request", err)
		}
		return &apiErr
	}

	result := h.mapper.MapEntityToResponse(*order)
	uri := fmt.Sprintf("/orders/%d", result.ID)
	ginCtx.Header("Location", uri)
	ginCtx.JSON(http.StatusCreated, result)
	return nil
}

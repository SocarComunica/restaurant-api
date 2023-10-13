package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/orders/core/entity"
	commonErrors "restaurant-api/src/api/errors"
	"restaurant-api/src/api/handler/order/contract"
)

type GetOrdersQueueUseCase interface {
	GetQueue(ctx *gin.Context) []entity.Order
}

type GetOrdersQueueHandler struct {
	useCase GetOrdersQueueUseCase
	mapper  Mapper
}

func NewGetOrdersQueueHandler(useCase GetOrdersQueueUseCase, mapper Mapper) GetOrdersQueueHandler {
	return GetOrdersQueueHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h GetOrdersQueueHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h GetOrdersQueueHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	orders := h.useCase.GetQueue(ginCtx)

	var result []contract.Order
	for _, order := range orders {
		result = append(result, h.mapper.MapEntityToResponse(order))
	}

	ginCtx.JSON(http.StatusOK, result)
	return nil
}

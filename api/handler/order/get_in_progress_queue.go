package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	commonErrors "restaurant-api/api/errors"
	"restaurant-api/api/handler/order/contract"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/orders/core/entity"
)

type GetInProgressOrdersQueueUseCase interface {
	GetInProgressOrdersQueue(ctx *gin.Context) []entity.Order
}

type GetInProgressOrdersQueueHandler struct {
	useCase GetInProgressOrdersQueueUseCase
	mapper  Mapper
}

func NewGetInProgressOrdersQueueHandler(useCase GetInProgressOrdersQueueUseCase, mapper Mapper) GetInProgressOrdersQueueHandler {
	return GetInProgressOrdersQueueHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h GetInProgressOrdersQueueHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h GetInProgressOrdersQueueHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	orders := h.useCase.GetInProgressOrdersQueue(ginCtx)

	var result []contract.Order
	for _, order := range orders {
		result = append(result, h.mapper.MapEntityToResponse(order))
	}

	ginCtx.JSON(http.StatusOK, result)
	return nil
}

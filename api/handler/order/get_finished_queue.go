package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	commonErrors "restaurant-api/api/errors"
	"restaurant-api/api/handler/order/contract"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/orders/core/entity"
)

type GetFinishedOrdersQueueUseCase interface {
	GetFinishedOrdersQueue(ctx *gin.Context) []entity.Order
}

type GetFinishedOrdersQueueHandler struct {
	useCase GetFinishedOrdersQueueUseCase
	mapper  Mapper
}

func NewGetFinishedOrdersQueueHandler(useCase GetFinishedOrdersQueueUseCase, mapper Mapper) GetFinishedOrdersQueueHandler {
	return GetFinishedOrdersQueueHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h GetFinishedOrdersQueueHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h GetFinishedOrdersQueueHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	orders := h.useCase.GetFinishedOrdersQueue(ginCtx)

	var result []contract.Order
	for _, order := range orders {
		result = append(result, h.mapper.MapEntityToResponse(order))
	}

	ginCtx.JSON(http.StatusOK, result)
	return nil
}

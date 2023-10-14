package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	commonErrors "restaurant-api/api/errors"
	"restaurant-api/api/handler/order/contract"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/orders/core/entity"
)

type GetQueuedOrdersQueueUseCase interface {
	GetQueuedOrdersQueue(ctx *gin.Context) []entity.Order
}

type GetQueuedOrdersQueueHandler struct {
	useCase GetQueuedOrdersQueueUseCase
	mapper  Mapper
}

func NewGetQueuedOrdersQueueHandler(useCase GetQueuedOrdersQueueUseCase, mapper Mapper) GetQueuedOrdersQueueHandler {
	return GetQueuedOrdersQueueHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h GetQueuedOrdersQueueHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h GetQueuedOrdersQueueHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	orders := h.useCase.GetQueuedOrdersQueue(ginCtx)

	var result []contract.Order
	for _, order := range orders {
		result = append(result, h.mapper.MapEntityToResponse(order))
	}

	ginCtx.JSON(http.StatusOK, result)
	return nil
}

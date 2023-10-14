package order

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	commonErrors "restaurant-api/api/errors"
	platformerrors "restaurant-api/internal/platform/error"
)

type UpdateOrdersQueuesUseCase interface {
	UpdateOrderQueuedToInProgress(ctx *gin.Context) error
	UpdateOrderInProgressToFinished(ctx *gin.Context) error
	UpdateOrderFinishedToDelivered(ctx *gin.Context) error
}

type UpdateOrdersQueuesHandler struct {
	useCase UpdateOrdersQueuesUseCase
}

func NewUpdateOrdersQueuesHandler(useCase UpdateOrdersQueuesUseCase) UpdateOrdersQueuesHandler {
	return UpdateOrdersQueuesHandler{
		useCase: useCase,
	}
}

func (h UpdateOrdersQueuesHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h UpdateOrdersQueuesHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	queue := ginCtx.Param("queue")

	var err error
	switch queue {
	case "queued":
		err = h.useCase.UpdateOrderQueuedToInProgress(ginCtx)
		break

	case "in-progress":
		err = h.useCase.UpdateOrderInProgressToFinished(ginCtx)
		break

	case "finished":
		err = h.useCase.UpdateOrderFinishedToDelivered(ginCtx)
		break

	default:
		message := fmt.Sprintf("there is no queued with parameter: %s", queue)
		apiErr := platformerrors.NewBadRequestAPIError(message, errors.New(message))
		return &apiErr
	}

	if err != nil {
		apiErr := platformerrors.NewInternalServerAPIError("an error has occurred while processing your request", err)
		return &apiErr
	}

	ginCtx.Status(http.StatusOK)
	return nil
}

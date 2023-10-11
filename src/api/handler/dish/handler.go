package dish

import (
	"github.com/gin-gonic/gin"
	"net/http"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/src/api/errors"
	"restaurant-api/src/api/handler/dish/contract"
)

type RandomDishUseCase interface {
	CreateRandomDish(ctx *gin.Context) (int, error)
}

type Handler struct {
	useCase RandomDishUseCase
}

func NewDishHandler(useCase RandomDishUseCase) Handler {
	return Handler{
		useCase: useCase,
	}
}

func (h Handler) Handler(ginCtx *gin.Context) {
	errors.ErrorWrapper(h.handler, ginCtx)
}

func (h Handler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	dishID, err := h.useCase.CreateRandomDish(ginCtx)
	if err != nil {
		apiErr := platformerrors.NewInternalServerAPIError("An error has occurred while processing your request", err)
		return &apiErr
	}

	ginCtx.JSON(http.StatusCreated, &contract.Response{DishId: dishID})
	return nil
}

package dish

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/internal/support/dishes/core/entity"
	commonErrors "restaurant-api/src/api/errors"
	"restaurant-api/src/api/handler/dish/mapper"
	"strconv"
)

const (
	NotFoundDishIDParamErrorMessage = "there is no resource available with this id"
)

type GetDishUseCase interface {
	GetDish(ctx *gin.Context, dishID int) (*entity.Dish, error)
}

type GetHandler struct {
	useCase GetDishUseCase
	mapper  Mapper
}

func NewGetHandler(useCase GetDishUseCase, mapper mapper.Mapper) GetHandler {
	return GetHandler{
		useCase: useCase,
		mapper:  mapper,
	}
}

func (h GetHandler) Handler(ginCtx *gin.Context) {
	commonErrors.ErrorWrapper(h.handler, ginCtx)
}

func (h GetHandler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	dishID, err := strconv.Atoi(ginCtx.Param("dishID"))
	if err != nil {
		apiErr := platformerrors.NewNotFoundAPIError(NotFoundDishIDParamErrorMessage, err)
		return &apiErr
	}

	dish, err := h.useCase.GetDish(ginCtx, dishID)
	if err != nil {
		var apiErr platformerrors.APIError
		if dish == nil {
			apiErr = platformerrors.NewNotFoundAPIError(NotFoundDishIDParamErrorMessage, errors.New(NotFoundDishIDParamErrorMessage))
		} else {
			apiErr = platformerrors.NewInternalServerAPIError("An error has occurred while processing your request", err)
		}
		return &apiErr
	}

	result := h.mapper.MapEntityToResponse(*dish)
	ginCtx.JSON(http.StatusOK, result)
	return nil
}

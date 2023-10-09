package dish

import (
	"github.com/gin-gonic/gin"
	platformerrors "restaurant-api/internal/platform/error"
	"restaurant-api/src/api/errors"
)

type Handler struct{}

func NewDishHandler() Handler {
	return Handler{}
}

func (h Handler) Handler(ginCtx *gin.Context) {
	errors.ErrorWrapper(h.handler, ginCtx)
}

func (h Handler) handler(ginCtx *gin.Context) *platformerrors.APIError {
	panic("IMPLEMENT THIS SHIT")
}

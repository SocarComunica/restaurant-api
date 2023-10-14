package errors

import (
	"github.com/gin-gonic/gin"
	platformerrors "restaurant-api/internal/platform/error"
)

type HandlerFunc func(c *gin.Context) *platformerrors.APIError

func ErrorWrapper(handlerFunc HandlerFunc, c *gin.Context) {
	if err := handlerFunc(c); err != nil {
		c.JSON(err.Status(), err)
	}
}

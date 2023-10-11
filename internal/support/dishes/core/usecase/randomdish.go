package usecase

import "github.com/gin-gonic/gin"

type RandomDishUseCase struct {
}

func NewRandomDishUseCase() RandomDishUseCase {
	return RandomDishUseCase{}
}

func (ruc RandomDishUseCase) CreateRandomDish(ctx *gin.Context) (int, error) {
	panic("Implement RandomDishUseCase")
}

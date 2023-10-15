package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"restaurant-api/internal/platform/restful"
	"restaurant-api/internal/support/warehouse/core/entity"
)

var validIngredients = map[string]string{
	"tomato":  "tomato",
	"lemon":   "lemon",
	"potato":  "potato",
	"rice":    "rice",
	"ketchup": "ketchup",
	"lettuce": "lettuce",
	"onion":   "onion",
	"cheese":  "cheese",
	"meat":    "meat",
	"chicken": "chicken",
}

type Repository interface {
	GetInventory(name string) (int, error)
	UseInventory(name string, amount int) error
	AddInventory(name string, amount int) error
}

type UseCase struct {
	repository Repository
	restClient restful.RestClient
}

func NewWarehouseUseCase(repository Repository, restClient restful.RestClient) UseCase {
	return UseCase{
		repository: repository,
		restClient: restClient,
	}
}

func (u *UseCase) GetIngredients(ctx *gin.Context, ingredients map[string]int) error {
	if ingredients == nil || len(ingredients) == 0 {
		return errors.New("there is no recipe, ingredients list is empty")
	}

	for name, amount := range ingredients {
		if err := u.geIngredient(ctx, name, amount); err != nil {
			return err
		}
	}

	return nil
}

func (u *UseCase) geIngredient(ctx *gin.Context, name string, amount int) error {
	if validIngredients[name] == "" {
		return errors.New(fmt.Sprintf("%s is not a valid ingredient", name))
	}

	currentAmount, err := u.repository.GetInventory(name)
	if err != nil {
		return err
	}

	if currentAmount < amount {
		response, err := u.restClient.Get(ctx, fmt.Sprintf("?ingredient=%s", name), nil)
		if err != nil {
			return err
		}

		var data entity.MarketPlaceResponse
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			return err
		}

		if data.Data[name] >= 1 {
			err = u.repository.AddInventory(name, data.Data[name])
			if err != nil {
				return err
			}
		} else {
			return u.geIngredient(ctx, name, amount)
		}
	}

	err = u.repository.UseInventory(name, amount)
	if err != nil {
		return err
	}
	return nil
}
